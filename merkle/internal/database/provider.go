package database

import (
	"github.com/Bedrock-Technology/VeMerkle/internal/database/psql"
	"github.com/sirupsen/logrus"
)

// GetMaxEpoch retrieves the maximum epoch from airdrop_data table
// Returns 0 if table is empty
func GetMaxEpoch() (uint64, error) {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return 0, err
	}

	var maxEpoch struct {
		MaxEpoch uint64
	}

	result := db.Model(&psql.AirdropData{}).
		Select("COALESCE(MAX(epoch), 0) as max_epoch").
		Scan(&maxEpoch)

	if result.Error != nil {
		return 0, result.Error
	}

	return maxEpoch.MaxEpoch, nil
}

// CheckEpochValidity validates if the given epoch is valid for new airdrop
// Returns true if:
// 1. Database is empty and epoch > 0 (first airdrop)
// 2. epoch = max_epoch + 1 (sequential airdrop)
func CheckEpochValidity(epoch uint64) (bool, error) {
	maxEpoch, err := GetMaxEpoch()
	if err != nil {
		return false, err
	}

	if maxEpoch == 0 {
		return epoch == 1, nil
	}

	return epoch == maxEpoch+1, nil
}

func CheckCurEpochValidity(epoch uint64) (bool, error) {
	maxEpoch, err := GetMaxEpoch()
	if err != nil {
		return false, err
	}
	if maxEpoch == 0 {
		return false, nil
	}
	return epoch == maxEpoch, nil
}

// GetAirdropByEpoch retrieves airdrop data for specific epoch
func GetAirdropByEpoch(epoch uint64) (*psql.AirdropData, error) {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return nil, err
	}

	var airdrop psql.AirdropData
	err = db.Where("epoch = ?", epoch).First(&airdrop).Error
	if err != nil {
		return nil, err
	}
	return &airdrop, nil
}

// BatchCreateAirdropData creates multiple airdrop records in database efficiently
func BatchCreateAirdropData(records []*psql.AirdropData) error {
	if len(records) == 0 {
		return nil
	}

	db, err := GetDBConnection("postgres")
	if err != nil {
		return err
	}

	// Use transaction for batch insertion
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Set larger batch size for better performance
	batchSize := 1000
	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}
		logrus.WithFields(logrus.Fields{
			"batch_start": i,
			"batch_end":   end,
			"batch_size":  end - i,
		}).Info("Processing batch")
		if err := tx.CreateInBatches(records[i:end], batchSize).Error; err != nil {
			tx.Rollback()
			return err
		}
		logrus.WithFields(logrus.Fields{
			"batch_start": i,
			"batch_end":   end,
		}).Info("Batch processed successfully")
	}

	return tx.Commit().Error
}

func GetUsersByEpoch(epoch uint64) ([]string, error) {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return nil, err
	}

	var users []string
	err = db.Model(&psql.AirdropData{}).Where("epoch = ?", epoch).Pluck("address", &users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateClaimedStatus(epoch uint64, users []string, claimedStatus []bool) error {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return err
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for i, user := range users {
		err := tx.Model(&psql.AirdropData{}).Where("epoch = ? AND address = ?", epoch, user).Update("claimed", claimedStatus[i]).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetAllAirdropDataByEpoch retrieves all airdrop data for a specific epoch
func GetAllAirdropDataByEpoch(epoch uint64) ([]*psql.AirdropData, error) {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return nil, err
	}

	var records []*psql.AirdropData
	err = db.Where("epoch = ?", epoch).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetClaimedAirdropDataByEpoch retrieves airdrop data for a specific epoch based on claimed status
func GetClaimedAirdropDataByEpoch(epoch uint64, claimed bool) ([]*psql.AirdropData, error) {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return nil, err
	}

	var records []*psql.AirdropData
	err = db.Where("epoch = ? AND claimed = ?", epoch, claimed).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

// DeleteAirdropDataByEpoch deletes all airdrop data for a specific epoch
func DeleteAirdropDataByEpoch(epoch uint64) error {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return err
	}

	// Use transaction for batch deletion
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Delete all records for the specified epoch
	if err := tx.Where("epoch = ?", epoch).Delete(&psql.AirdropData{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// CheckEpochExists checks if any records exist for the given epoch
// Returns true if records exist, false otherwise
func CheckEpochExists(epoch uint64) (bool, error) {
	db, err := GetDBConnection("postgres")
	if err != nil {
		return false, err
	}

	var count int64
	err = db.Model(&psql.AirdropData{}).
		Where("epoch = ?", epoch).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
