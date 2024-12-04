package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var suffix string
	var num int

	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "generate-address",
		Short: "Generate Tron addresses with a specific suffix",
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < num; i++ {
				privateKeyHex, trxAddress := GenerateNiceKey(suffix)
				fmt.Printf("%s\t%s\n", privateKeyHex, trxAddress)
			}
		},
	}

	// Add flags for suffix and number of addresses
	rootCmd.Flags().StringVarP(&suffix, "suffix", "s", "888", "The desired suffix for the Tron address")
	rootCmd.Flags().IntVarP(&num, "num", "n", 1, "The number of addresses to generate")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// GenerateKey generates a new private key and Tron address.
func GenerateKey() (privateKeyHex string, trxAddress string, err error) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}

	pubKey := privateKey.PubKey().ToECDSA()
	trxAddr := address.PubkeyToAddress(*pubKey)
	trxAddress = trxAddr.String()

	privateKeyHex = hex.EncodeToString(privateKey.Serialize())

	return privateKeyHex, trxAddress, nil
}

// GenerateNiceKey generates a Tron address with the specified suffix.
func GenerateNiceKey(suffix string) (privateKeyHex string, trxAddress string) {
	var err error
	for {
		privateKeyHex, trxAddress, err = GenerateKey()
		if err != nil {
			panic(err)
		}

		if len(trxAddress) >= len(suffix) && trxAddress[len(trxAddress)-len(suffix):] == suffix {
			return privateKeyHex, trxAddress
		}
	}
}
