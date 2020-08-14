package contract

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var managerKSPool, validatorKSPool RoundRobin

type KSItem struct {
	Key    string
	Secret string
}

var managerKS = []*KSItem{}
var validatorKS = []*KSItem{}

func shuffle(vals []*KSItem) []*KSItem {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]*KSItem, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func GetManagerKS() ([]byte, string) {
	item := managerKSPool.Next()
	return []byte(item.Key), item.Secret
}

func GetValidatorKS() ([]byte, string) {
	item := validatorKSPool.Next()
	return []byte(item.Key), item.Secret
}

func loadKSFromFile(path string) ([]*KSItem, error) {
	ks := []*KSItem{}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			continue
		}

		item := &KSItem{
			Secret: strings.TrimSpace(parts[0]),
			Key:    strings.TrimSpace(parts[1]),
		}

		fmt.Printf("%+v\n", item)

		ks = append(ks, item)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(ks) == 0 {
		return nil, errors.New("empty ks list")
	}

	return ks, nil
}

func LoadKSFromFiles(managerKSPath, validatorKSPath string) error {
	mks, err := loadKSFromFile(managerKSPath)
	if err != nil {
		return err
	}

	vks, err := loadKSFromFile(validatorKSPath)
	if err != nil {
		return err
	}

	for _, item := range mks {
		managerKS = append(managerKS, item)
	}

	for _, item := range vks {
		validatorKS = append(validatorKS, item)
	}

	managerKS = shuffle(managerKS)
	validatorKS = shuffle(validatorKS)

	managerKSPool, err = NewRoundRobin(managerKS)
	if err != nil {
		return err
	}

	validatorKSPool, err = NewRoundRobin(validatorKS)
	if err != nil {
		return err
	}

	return nil
}
