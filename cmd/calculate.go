/*
 * Project: CRC32 Checksum Tools
 * Filename: /cmd/calculate.go
 * Created Date: Tuesday September 19th 2023 18:39:08 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Wednesday September 20th 2023 13:37:57 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package cmd

import (
	"fmt"

	"github.com/berrypay/appsec"
	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate CRC32 checksum",
	Long:  `Calculate CRC32 checksum for a given list of arguments using the specified polynomial.`,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := cmd.Flags().GetString("polynomial")
		switch p {
		case "koopman":
			doKoopman(args)
		case "ieee":
			doIEEE(args)
		default:
			doCastagnoli(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doCastagnoli(args []string) {
	crcSrc := make([]byte, len(args))
	for _, v := range args {
		crcSrc = append(crcSrc, []byte(v)...)
	}
	fmt.Printf("CRC32 String: %s\n", crcSrc)
	fmt.Printf("CRC32 Byte: %d\n", crcSrc)
	fmt.Printf("CRC32 Hex: %x\n", crcSrc)
	crc := appsec.Crc32Castagnoli(crcSrc)
	fmt.Printf("%d\n", crc)
}

func doKoopman(args []string) {
	// TODO: Implement Koopman CRC32 calculation
}

func doIEEE(args []string) {
	// TODO: Implement IEEE CRC32 calculation
}
