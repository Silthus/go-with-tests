package maps

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DictionaryTestSuite struct {
	suite.Suite
	dictionary Dictionary
}

func (s *DictionaryTestSuite) SetupTest() {
	s.dictionary = Dictionary{
		"test": "just a test",
	}
}

func (s *DictionaryTestSuite) TestSearch() {
	s.Run("known word", func() {
		s.assertDefinition("test", "just a test")
	})
	s.Run("unknown word", func() {
		_, err := s.dictionary.Search("unknown")
		s.ErrorIs(err, ErrNotFound)
	})
}

func (s *DictionaryTestSuite) TestAdd() {
	s.Run("new definition", func() {
		err := s.dictionary.Add("foo", "bar")
		s.NoError(err)
		s.assertDefinition("foo", "bar")
	})
	s.Run("existing definition throws error", func() {
		err := s.dictionary.Add("test", "should fail")
		s.ErrorIs(err, ErrWordExists)
		s.assertDefinition("test", "just a test")
	})
}

func (s *DictionaryTestSuite) TestUpdate() {
	s.Run("new definition is added", func() {
		s.dictionary.Update("foo", "bar")
		s.assertDefinition("foo", "bar")
	})
	s.Run("existing definition is updated", func() {
		s.dictionary.Update("test", "updated test")
		s.assertDefinition("test", "updated test")
	})
}

func (s *DictionaryTestSuite) assertDefinition(word, expectedResult string) {
	s.T().Helper()
	result, err := s.dictionary.Search(word)
	s.NoError(err)
	s.Equal(expectedResult, result)
}

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(DictionaryTestSuite))
}
