package cmd

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/Bedrock-Technology/VeMerkle/internal/config"
	"github.com/Bedrock-Technology/VeMerkle/internal/logger"
	"github.com/Bedrock-Technology/VeMerkle/internal/proto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	port     int
	merkleDB *merkleTree
)

type merkleTree struct {
	tree      *smt.StandardTree
	addresses map[string]int
	amounts   map[string]*big.Int
}

// @title Airdrop Merkle API
// @version 1.0
// @description API for verifying Merkle proofs of airdrop data
// @host localhost:8080
// @BasePath /api/v1
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the Merkle proof verification API server",
	Long: `Start a HTTP server that provides API endpoints for Merkle proof verification.
    The Merkle tree data should be imported via the /merkle/import API endpoint.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load configuration
		configPath := filepath.Join("configs", "airdrop.yaml")
		appConfig, err := config.LoadConfig(configPath)
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		// Initialize logger with loaded configuration
		if err := logger.InitLogger(appConfig.Logger); err != nil {
			return fmt.Errorf("failed to initialize logger: %v", err)
		}

		// Setup Gin router
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Recovery())
		r.Use(loggerMiddleware())
		r.Use(CORSMiddleware())
		setupRouter(r)
		setupSwagger(r, appConfig.DocAuth)

		// Start server
		addr := fmt.Sprintf(":%d", port)
		logrus.WithField("port", port).Info("Server starting")
		return r.Run(addr)
	},
}

func setupRouter(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		// API endpoints
		api.POST("/merkle/proof", getMerkleProof)
		api.POST("/merkle/import", importMerkleTreeHandler)
		api.GET("/merkle/root", getMerkleRoot)
	}
}

func setupSwagger(e *gin.Engine, docAuth map[string]string) {
	auth := gin.Accounts(docAuth)
	e.GET("/docs/*any", gin.BasicAuth(auth), ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.StaticFile("/swagger/doc.json", "./docs/swagger.json")
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&port, "port", 8080, "Port to run the server on")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// loggerMiddleware creates a Gin middleware for logging HTTP requests
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Log request details
		logrus.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
			"latency":    c.Writer.Header().Get("X-Response-Time"),
		}).Info("HTTP request")
	}
}

// getMerkleProof handles the Merkle proof API endpoint
// @Summary Get Merkle proof for address
// @Description Get Merkle proof, amount and root hash for a given address
// @Tags merkle
// @Accept json
// @Produce json
// @Param address body string true "Ethereum address"
// @Success 200 {object} map[string]interface{} "Returns address, amount, proof array and root hash"
// @Failure 404 {object} map[string]string "Returns error message when address not found"
// @Router /merkle/proof [post]
func getMerkleProof(c *gin.Context) {
	var request struct {
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		proto.ErrorMsg(c, "failed to bind JSON")
		return
	}

	address := strings.ToLower(request.Address)

	// Check if address exists
	index, exists := merkleDB.addresses[address]
	if !exists {
		logrus.WithField("address", request.Address).Warn("Address not found")
		proto.ErrorMsg(c, "address not found")
		return
	}
	logrus.WithFields(logrus.Fields{
		"query address": address,
		"in array indx": index,
	}).Info("Reading merkle db")

	// Get amount and proof
	amount := merkleDB.amounts[address]
	leaf := []interface{}{
		smt.SolAddress(address),
		smt.SolNumber(amount.String()),
	}

	proof, err := merkleDB.tree.GetProof(leaf)
	if err != nil {
		logrus.WithError(err).Error("Failed to get Merkle proof")
		proto.ErrorMsg(c, "failed to get Merkle proof")
		return
	}

	// Verify the proof
	verify, err := merkleDB.tree.Verify(proof, leaf)
	if err != nil {
		logrus.WithError(err).Error("Failed to verify Merkle proof")
		proto.ErrorMsg(c, "failed to verify Merkle proof")
		return
	}
	if !verify {
		logrus.WithFields(logrus.Fields{
			"address": address,
			"amount":  amount.String(),
		}).Error("Invalid Merkle proof generated")
		proto.ErrorMsg(c, "invalid merkle proof")
		return
	}

	// Convert proof to hex strings
	hexProof := make([]string, len(proof))
	for i, p := range proof {
		hexProof[i] = hex.EncodeToString([]byte(p))
	}

	logrus.WithFields(logrus.Fields{
		"address": address,
		"amount":  amount.String(),
	}).Info("Generated Merkle proof")

	treeRoot := hexutil.Encode(merkleDB.tree.GetRoot())
	proto.SuccessMsg(c, http.StatusOK, "Merkle proof generated successfully", gin.H{
		"address": address,
		"amount":  amount.String(),
		"proof":   hexProof,
		"root":    treeRoot,
	})
}

// importMerkleTree reads a CSV file and updates the merkleDB with new data
func importMerkleTree(csvFile string) error {
	// Open CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %v", err)
	}

	logrus.WithField("records length", len(records)).Info("Reading CSV file")
	// Skip header row
	if len(records) < 2 {
		return fmt.Errorf("CSV file is empty or contains only header")
	}
	records = records[1:]

	// Prepare new data for Merkle tree
	addresses := make(map[string]int)
	amounts := make(map[string]*big.Int)
	values := [][]interface{}{}
	for i, record := range records {
		if len(record) != 2 {
			return fmt.Errorf("invalid CSV record format at line %d", i+2)
		}
		address := strings.ToLower(record[0])
		amount, ok := new(big.Int).SetString(record[1], 10)
		if !ok {
			return fmt.Errorf("invalid amount format at line %d", i+2)
		}
		logrus.WithFields(logrus.Fields{
			"address": address,
			"amount":  amount.String(),
			"index":   i,
		}).Info("Reading CSV file")
		addresses[address] = i
		amounts[address] = amount
		values = append(values, []interface{}{
			smt.SolAddress(address),
			smt.SolNumber(record[1]),
		})
	}

	leafEncodings := []string{
		smt.SOL_ADDRESS,
		smt.SOL_UINT256,
	}
	if values == nil {
		return fmt.Errorf("values are nil")
	}

	tree, err := smt.Of(values, leafEncodings)
	merkleDB = &merkleTree{
		tree:      tree,
		addresses: addresses,
		amounts:   amounts,
	}

	return nil
}

// @Summary Import Merkle tree data from CSV
// @Description Import address and amount data from a CSV file to update the Merkle tree
// @Tags merkle
// @Accept json
// @Produce json
// @Param csvFile body string true "Full path to the CSV file containing address and amount"
// @Success 200 {object} map[string]string "Returns success message"
// @Failure 400 {object} map[string]string "Returns error message when CSV file is invalid"
// @Router /merkle/import [post]
func importMerkleTreeHandler(c *gin.Context) {
	var request struct {
		CsvFile string `json:"csvFile"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		proto.ErrorMsg(c, "failed to bind JSON")
		return
	}

	// Import Merkle tree data
	if err := importMerkleTree(request.CsvFile); err != nil {
		logrus.WithError(err).Error("Failed to import Merkle tree")
		proto.ErrorMsg(c, err.Error())
		return
	}

	proto.SuccessMsg(c, http.StatusOK, "Merkle tree imported successfully", nil)
}

// @Summary Get Merkle tree root
// @Description Retrieve the root hash of the Merkle tree
// @Tags merkle
// @Produce json
// @Success 200 {object} map[string]string "Returns the root hash of the Merkle tree"
// @Failure 404 {object} map[string]string "Returns error message when Merkle tree is not initialized"
// @Router /merkle/root [get]
func getMerkleRoot(c *gin.Context) {
	if merkleDB == nil || merkleDB.tree == nil {
		proto.ErrorMsg(c, "Merkle tree not initialized")
		return
	}
	treeRoot := hexutil.Encode(merkleDB.tree.GetRoot())
	logrus.WithField("root", treeRoot).Info("Retrieved Merkle tree root")
	proto.SuccessMsg(c, http.StatusOK, "Merkle tree root retrieved successfully", gin.H{
		"root": treeRoot,
	})
}
