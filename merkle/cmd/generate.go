package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	// Parameters for normalize-csv command
	inputFile     string
	outputFile    string
	airdropAmount string
	minAmount     string

	// Parameters for import-merkle-csv command
	rpcEndpoint string
	merkleFile  string
	epoch       uint64
	persist     bool

	// Parameters for export-csv command
	exportEpoch   uint64
	exportClaimed int64

	// Parameters for import-airdrop command
	importEpoch   uint64
	importCsvFile string

	// Parameters for update-claim command
	updateClaimEpoch uint64

	// Parameters for delete command
	deleteEpoch uint64
)

// importAirdropCmd represents the import-airdrop subcommand
var importAirdropCmd = &cobra.Command{
	Use:   "import-airdrop",
	Short: "Import airdrop data from CSV file",
	Long: `Import airdrop data from a CSV file for a specific epoch.
	    The CSV file should contain address and amount columns.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Prepare request body
		requestBody := map[string]interface{}{
			"epoch":   importEpoch,
			"csvFile": importCsvFile,
		}

		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}

		// Create HTTP request
		url := fmt.Sprintf("%s/api/v1/merkle/import_airdrop", rpcEndpoint)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("API request failed: %s", string(body))
		}

		fmt.Printf("Successfully imported airdrop data from %s\n", importCsvFile)
		return nil
	},
}

// generateCmd represents the parent command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate and process airdrop data",
	Long: `Generate and process airdrop data in CSV format.
	Supports generating random data and normalizing existing CSV files.`,
}

// normalizeCSVCmd represents the normalize-csv subcommand
var normalizeCSVCmd = &cobra.Command{
	Use:   "normalize-csv",
	Short: "Normalize airdrop CSV file format",
	Long: `Normalize an existing CSV file to ensure it meets the required format.
	Validates addresses and amounts, normalizes amounts based on total supply,
	removes entries below minimum amount, and removes duplicates.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(inputFile)
		if err != nil {
			return fmt.Errorf("failed to open input file: %v", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read CSV file: %v", err)
		}

		if len(records) < 2 {
			return fmt.Errorf("CSV file is empty or contains only header")
		}

		// Parse airdrop amount and minimum amount
		totalAirdrop, ok := new(big.Int).SetString(airdropAmount, 10)
		if !ok {
			return fmt.Errorf("invalid airdrop amount format")
		}

		minAmountValue, ok := new(big.Int).SetString(minAmount, 10)
		if !ok {
			return fmt.Errorf("invalid minimum amount format")
		}

		// Calculate total amount and validate data
		totalAmount := new(big.Int)
		normalized := make(map[string]*big.Int)

		for i, record := range records[1:] {
			if len(record) != 2 {
				return fmt.Errorf("invalid CSV record format at line %d", i+2)
			}

			address := strings.TrimSpace(record[0])
			if !isValidEthAddress(address) {
				return fmt.Errorf("invalid ethereum address at line %d: %s", i+2, address)
			}
			amount, ok := new(big.Int).SetString(strings.TrimSpace(record[1]), 10)
			if !ok {
				return fmt.Errorf("invalid amount format at line %d", i+2)
			}

			// Merge amounts for duplicate addresses
			if existing, exists := normalized[address]; exists {
				normalized[address] = new(big.Int).Add(existing, amount)
			} else {
				normalized[address] = amount
			}
			totalAmount = new(big.Int).Add(totalAmount, amount)
		}

		// Normalize amounts and filter by minimum amount
		filteredData := make(map[string]*big.Int)
		for address, amount := range normalized {
			// Calculate normalized amount: (amount / totalAmount) * airdropAmount
			normalizedAmount := new(big.Int).Mul(amount, totalAirdrop)
			normalizedAmount = normalizedAmount.Div(normalizedAmount, totalAmount)

			// Filter out amounts below minimum
			if normalizedAmount.Cmp(minAmountValue) >= 0 {
				filteredData[address] = normalizedAmount
			}
		}

		// Create output file
		if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}

		outFile, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer outFile.Close()

		writer := csv.NewWriter(outFile)
		defer writer.Flush()

		if err := writer.Write([]string{"address", "amount"}); err != nil {
			return fmt.Errorf("failed to write CSV header: %v", err)
		}

		for address, amount := range filteredData {
			fmt.Printf("Normalized: %s -> %s\n", address, amount.String())
			if err := writer.Write([]string{address, amount.String()}); err != nil {
				return fmt.Errorf("failed to write CSV record: %v", err)
			}
		}

		fmt.Printf("Normalized %d addresses in %s (filtered from %d original addresses)\n",
			len(filteredData), outputFile, len(normalized))
		return nil
	},
}

// importMerkleCSVCmd represents the import-merkle-csv subcommand
var importMerkleCSVCmd = &cobra.Command{
	Use:   "import-merkle-csv",
	Short: "Import CSV file to Merkle tree via RPC",
	Long: `Import a CSV file to update the Merkle tree by calling the RPC endpoint.
	The CSV file should contain address and amount columns.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate RPC endpoint
		if !strings.HasPrefix(rpcEndpoint, "http://") && !strings.HasPrefix(rpcEndpoint, "https://") {
			return fmt.Errorf("invalid RPC endpoint format, must start with http:// or https://")
		}

		// Prepare request URL
		importURL := fmt.Sprintf("%s/api/v1/merkle/import", strings.TrimSuffix(rpcEndpoint, "/"))

		// Prepare request body
		reqBody := map[string]interface{}{
			"csvFile": merkleFile,
			"epoch":   epoch,
			"persist": persist,
		}
		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}

		// Create HTTP request
		req, err := http.NewRequest("POST", importURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %v", err)
		}

		// Check response status
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("import failed: %s", string(body))
		}

		fmt.Printf("Successfully imported Merkle tree from %s\n", merkleFile)
		return nil
	},
}

// exportCSVCmd represents the export-csv subcommand
var exportCSVCmd = &cobra.Command{
	Use:   "export-csv",
	Short: "Export airdrop data to CSV file",
	Long: `Export airdrop data for a specific epoch to a CSV file.
	Supports exporting all data, claimed data, or unclaimed data based on the claimed parameter.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Prepare request body
		requestBody := map[string]interface{}{
			"epoch":   exportEpoch,
			"csvFile": outputFile,
			"claimed": exportClaimed,
		}

		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}

		// Create HTTP request
		url := fmt.Sprintf("%s/api/v1/merkle/export_airdrop", rpcEndpoint)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("API request failed: %s", string(body))
		}

		fmt.Printf("Successfully exported airdrop data to %s\n", outputFile)
		return nil
	},
}

// updateClaimCmd represents the update-claim subcommand
var updateClaimCmd = &cobra.Command{
	Use:   "update-claim",
	Short: "Update claim status for airdrop data",
	Long: `Update claim status for airdrop data of a specific epoch by calling the API endpoint.
    This will sync the on-chain claim status with the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Prepare request body
		requestBody := map[string]interface{}{
			"epoch": updateClaimEpoch,
		}

		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}

		// Create HTTP request
		url := fmt.Sprintf("%s/api/v1/merkle/claimed", rpcEndpoint)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("API request failed: %s", string(body))
		}

		fmt.Printf("Successfully updated claim status for epoch %d\n", updateClaimEpoch)
		return nil
	},
}

// deleteCmd represents the delete subcommand
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete airdrop data for specific epoch",
	Long: `Delete all airdrop data for a specific epoch from the database.
    This operation cannot be undone.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Prepare request body
		requestBody := map[string]interface{}{
			"epoch": deleteEpoch,
		}

		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}

		// Create HTTP request
		url := fmt.Sprintf("%s/api/v1/merkle/delete_airdrop", rpcEndpoint)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("API request failed: %s", string(body))
		}

		fmt.Printf("Successfully deleted airdrop data for epoch %d\n", deleteEpoch)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(normalizeCSVCmd)

	// Parameters for normalize-csv command
	normalizeCSVCmd.Flags().StringVar(&inputFile, "in", "", "Input CSV file path")
	normalizeCSVCmd.Flags().StringVar(&outputFile, "out", "normalized.csv", "Output CSV file path")
	normalizeCSVCmd.Flags().StringVar(&airdropAmount, "amount", "", "Total airdrop amount for normalization")
	normalizeCSVCmd.Flags().StringVar(&minAmount, "min", "0", "Minimum amount threshold (addresses with normalized amount below this will be filtered out)")

	normalizeCSVCmd.MarkFlagRequired("in")
	normalizeCSVCmd.MarkFlagRequired("amount")

	// Add import-merkle-csv command
	generateCmd.AddCommand(importMerkleCSVCmd)

	// Add flags for import-merkle-csv command
	importMerkleCSVCmd.Flags().StringVar(&rpcEndpoint, "rpc", "http://localhost:8080", "RPC endpoint URL")
	importMerkleCSVCmd.Flags().StringVar(&merkleFile, "file", "", "CSV file path")
	importMerkleCSVCmd.Flags().Uint64Var(&epoch, "epoch", 0, "Airdrop epoch number")
	importMerkleCSVCmd.Flags().BoolVar(&persist, "persist", false, "Persist data to database")

	importMerkleCSVCmd.MarkFlagRequired("file")
	importMerkleCSVCmd.MarkFlagRequired("epoch")
	importMerkleCSVCmd.MarkFlagRequired("persist")

	// Add export-csv command
	generateCmd.AddCommand(exportCSVCmd)

	// Add flags for export-csv command
	exportCSVCmd.Flags().StringVar(&rpcEndpoint, "rpc", "http://localhost:8080", "RPC endpoint URL")
	exportCSVCmd.Flags().StringVar(&outputFile, "out", "export.csv", "Output CSV file path")
	exportCSVCmd.Flags().Uint64Var(&exportEpoch, "epoch", 0, "Airdrop epoch number")
	exportCSVCmd.Flags().Int64Var(&exportClaimed, "claimed", 0, "Export type (0: all, 1: claimed, 2: unclaimed)")

	exportCSVCmd.MarkFlagRequired("epoch")

	// Add import-airdrop command
	generateCmd.AddCommand(importAirdropCmd)

	// Add flags for import-airdrop command
	importAirdropCmd.Flags().StringVar(&rpcEndpoint, "rpc", "http://localhost:8080", "RPC endpoint URL")
	importAirdropCmd.Flags().StringVar(&importCsvFile, "file", "", "CSV file path")
	importAirdropCmd.Flags().Uint64Var(&importEpoch, "epoch", 0, "Airdrop epoch number")

	importAirdropCmd.MarkFlagRequired("file")
	importAirdropCmd.MarkFlagRequired("epoch")

	// Add update-claim command
	generateCmd.AddCommand(updateClaimCmd)

	// Add flags for update-claim command
	updateClaimCmd.Flags().StringVar(&rpcEndpoint, "rpc", "http://localhost:8080", "RPC endpoint URL")
	updateClaimCmd.Flags().Uint64Var(&updateClaimEpoch, "epoch", 0, "Airdrop epoch number")

	updateClaimCmd.MarkFlagRequired("epoch")

	// Add delete command
	generateCmd.AddCommand(deleteCmd)

	// Add flags for delete command
	deleteCmd.Flags().StringVar(&rpcEndpoint, "rpc", "http://localhost:8080", "RPC endpoint URL")
	deleteCmd.Flags().Uint64Var(&deleteEpoch, "epoch", 0, "Airdrop epoch number")

	deleteCmd.MarkFlagRequired("epoch")
}

// Helper function: Validate Ethereum address format with EIP-55 checksum
func isValidEthAddress(address string) bool {
	if !common.IsHexAddress(address) {
		return false
	}

	// Check if the address follows EIP-55 checksum format
	return address == common.HexToAddress(address).Hex()
}

func isValidAddress(address string) bool {
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	address = strings.TrimPrefix(address, "0x")
	return len(address) == 40 && isHexString(address)
}

func isHexString(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
