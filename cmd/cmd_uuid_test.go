package cmd_test

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"uuid/cmd"
	"uuid/internal/assert"

	"github.com/gofrs/uuid/v5"
)

func TestV1Cmd(t *testing.T) {
	t.Run(`use is "v1"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V1Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "v1")
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V1Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 1)
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V1Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), number)
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, actual, 1)
		}
	})
}

func TestV3Cmd(t *testing.T) {
	t.Run(`use is "v3 [value]"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V3Cmd(uuid.NamespaceDNS)
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "v3 [value]")
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 3)
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), number)
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, actual, 3)
		}
	})

	t.Run("return error on invalid namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = "invalid"
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.Error(t, err)
	})

	t.Run("generate uuid with namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = uuid.Must(uuid.NewV4()).String()
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 3)
	})
}

func TestV4Cmd(t *testing.T) {
	t.Run(`use is "v4"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V4Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "v4")
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V4Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 4)
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V4Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), number)
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, actual, 4)
		}
	})
}

func TestV5Cmd(t *testing.T) {
	t.Run(`use is "v5 [value]"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V5Cmd(uuid.NamespaceDNS)
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "v5 [value]")
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 5)
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), number)
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, actual, 5)
		}
	})

	t.Run("return error on invalid namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = "invalid"
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.Error(t, err)
	})

	t.Run("generate uuid with namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = uuid.Must(uuid.NewV4()).String()
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 5)
	})
}

func TestV6Cmd(t *testing.T) {
	t.Run(`use is "v6"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V6Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "v6")
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V6Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 6)
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V6Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), number)
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, actual, 6)
		}
	})
}

func TestV7Cmd(t *testing.T) {
	t.Run(`use is "v7"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V7Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "v7")
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 7)
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), number)
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, actual, 7)
		}
	})

	t.Run("return error on invalid epoch", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			epoch      = "invalid"
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagEpoch, epoch)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.Error(t, err)
	})

	t.Run("generate uuid with epoch", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			epoch      = time.Now().Format(time.RFC3339Nano)
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagEpoch, epoch)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, string(actual), 7)
	})
}

func TestNullCmd(t *testing.T) {
	t.Run(`use is "null"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.NullCmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, actual, "null")
	})

	t.Run("generate null UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.NullCmd()
		)
		sut.SetOut(writerMock)

		// act
		sut.Run(sut, nil)

		// assert
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.Equal(t, string(actual), "00000000-0000-0000-0000-000000000000")
	})
}
