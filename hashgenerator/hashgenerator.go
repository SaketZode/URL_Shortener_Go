package hashgenerator

// Interface for hash generator.
type HashGenerator interface {
	GenerateHash(string) string
}
