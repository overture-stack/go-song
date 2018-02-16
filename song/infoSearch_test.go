/*
 *     Copyright (C) 2018  Ontario Institute for Cancer Research
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package song

import (
	"fmt"
)

func ExampleSearch() {
	x := map[string]string{"search1": "one", "search2": "two"}
	y := createInfoSearchJSON(true, x)
	fmt.Printf("y=%s\n", y)

	x = map[string]string{"a": "1", "b": "2", "c": "3"}
	y = createInfoSearchJSON(false, x)
	fmt.Printf("y=%s", y)

	//Output:
	//y={"includeInfo":true,"searchTerms":[{"key":"search1","value":"one"},{"key":"search2","value":"two"}]}
	//y={"includeInfo":false,"searchTerms":[{"key":"a","value":"1"},{"key":"b","value":"2"},{"key":"c","value":"3"}]}
}
