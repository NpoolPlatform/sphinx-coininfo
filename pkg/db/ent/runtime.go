// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/schema"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/transaction"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	coininfoFields := schema.CoinInfo{}.Fields()
	_ = coininfoFields
	// coininfoDescName is the schema descriptor for name field.
	coininfoDescName := coininfoFields[2].Descriptor()
	// coininfo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	coininfo.NameValidator = coininfoDescName.Validators[0].(func(string) error)
	// coininfoDescUnit is the schema descriptor for unit field.
	coininfoDescUnit := coininfoFields[3].Descriptor()
	// coininfo.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	coininfo.UnitValidator = coininfoDescUnit.Validators[0].(func(string) error)
	// coininfoDescIsPresale is the schema descriptor for is_presale field.
	coininfoDescIsPresale := coininfoFields[4].Descriptor()
	// coininfo.DefaultIsPresale holds the default value on creation for the is_presale field.
	coininfo.DefaultIsPresale = coininfoDescIsPresale.Default.(bool)
	// coininfoDescLogoImage is the schema descriptor for logo_image field.
	coininfoDescLogoImage := coininfoFields[5].Descriptor()
	// coininfo.DefaultLogoImage holds the default value on creation for the logo_image field.
	coininfo.DefaultLogoImage = coininfoDescLogoImage.Default.(string)
	// coininfoDescID is the schema descriptor for id field.
	coininfoDescID := coininfoFields[0].Descriptor()
	// coininfo.DefaultID holds the default value on creation for the id field.
	coininfo.DefaultID = coininfoDescID.Default.(func() uuid.UUID)
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescIsApproved is the schema descriptor for is_approved field.
	reviewDescIsApproved := reviewFields[1].Descriptor()
	// review.DefaultIsApproved holds the default value on creation for the is_approved field.
	review.DefaultIsApproved = reviewDescIsApproved.Default.(bool)
	// reviewDescOperatorNote is the schema descriptor for operator_note field.
	reviewDescOperatorNote := reviewFields[2].Descriptor()
	// review.OperatorNoteValidator is a validator for the "operator_note" field. It is called by the builders before save.
	review.OperatorNoteValidator = reviewDescOperatorNote.Validators[0].(func(string) error)
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescAddressFrom is the schema descriptor for address_from field.
	transactionDescAddressFrom := transactionFields[3].Descriptor()
	// transaction.AddressFromValidator is a validator for the "address_from" field. It is called by the builders before save.
	transaction.AddressFromValidator = func() func(string) error {
		validators := transactionDescAddressFrom.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address_from string) error {
			for _, fn := range fns {
				if err := fn(address_from); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// transactionDescAddressTo is the schema descriptor for address_to field.
	transactionDescAddressTo := transactionFields[4].Descriptor()
	// transaction.AddressToValidator is a validator for the "address_to" field. It is called by the builders before save.
	transaction.AddressToValidator = func() func(string) error {
		validators := transactionDescAddressTo.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address_to string) error {
			for _, fn := range fns {
				if err := fn(address_to); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// transactionDescNeedManualReview is the schema descriptor for need_manual_review field.
	transactionDescNeedManualReview := transactionFields[5].Descriptor()
	// transaction.DefaultNeedManualReview holds the default value on creation for the need_manual_review field.
	transaction.DefaultNeedManualReview = transactionDescNeedManualReview.Default.(bool)
	// transactionDescTransactionIDInsite is the schema descriptor for transaction_id_insite field.
	transactionDescTransactionIDInsite := transactionFields[7].Descriptor()
	// transaction.TransactionIDInsiteValidator is a validator for the "transaction_id_insite" field. It is called by the builders before save.
	transaction.TransactionIDInsiteValidator = transactionDescTransactionIDInsite.Validators[0].(func(string) error)
	// transactionDescTransactionIDChain is the schema descriptor for transaction_id_chain field.
	transactionDescTransactionIDChain := transactionFields[8].Descriptor()
	// transaction.TransactionIDChainValidator is a validator for the "transaction_id_chain" field. It is called by the builders before save.
	transaction.TransactionIDChainValidator = transactionDescTransactionIDChain.Validators[0].(func(string) error)
	// transactionDescMutex is the schema descriptor for mutex field.
	transactionDescMutex := transactionFields[10].Descriptor()
	// transaction.DefaultMutex holds the default value on creation for the mutex field.
	transaction.DefaultMutex = transactionDescMutex.Default.(bool)
	// transactionDescSignatureUser is the schema descriptor for signature_user field.
	transactionDescSignatureUser := transactionFields[11].Descriptor()
	// transaction.SignatureUserValidator is a validator for the "signature_user" field. It is called by the builders before save.
	transaction.SignatureUserValidator = transactionDescSignatureUser.Validators[0].(func(string) error)
	// transactionDescSignaturePlatform is the schema descriptor for signature_platform field.
	transactionDescSignaturePlatform := transactionFields[12].Descriptor()
	// transaction.SignaturePlatformValidator is a validator for the "signature_platform" field. It is called by the builders before save.
	transaction.SignaturePlatformValidator = transactionDescSignaturePlatform.Validators[0].(func(string) error)
}
