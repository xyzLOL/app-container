// Copyright [2016] [Cuiting Shi ]
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
// http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// 
package order

import (
	"cbdforum/app-container/chaincode/src/common"
)

type RequestType interface {
	GetType() string
}

func CheckArguments(args []string, expectedNum int) error {
	var err error
	if len(args) != expectedNum {
		return NewOrderErrorMessage(ERROR_ARGUMENTS, "Incorrect number, expecting %d", expectedNum)
	}
	//input sanitation
	err = common.SanitizeArguments(args)
	if err != nil {
		return NewOrderErrorMessage(ERROR_ARGUMENTS, "Invalid arguments: %v", err)
	}
	return nil
}
