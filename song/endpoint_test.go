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
import ("fmt" 
	"net/url")

func createEndpoint(address string) Endpoint {
	a, err := url.Parse(address)
	if err != nil { panic(err) } 
	return Endpoint{a}
}

func display(address url.URL) {
   fmt.Println(address.String())
}

func ExampleUpload() {
    e := createEndpoint("http://test.com")
    studyId := "ABC123"
    isAsync := false

    x := e.Upload(studyId, isAsync)

    isAsync = true 
    y := e.Upload(studyId, isAsync)

    display(x)
    display(y)

    // Output: 
    // http://test.com/upload/ABC123 
    // http://test.com/upload/ABC123/async
}

func ExampleGetStatus() {
    e := createEndpoint("http://www.mrap.org")
    studyId := "XYZ234"
    uploadId := "UP-AB2345"

    x := e.GetStatus(studyId, uploadId)
    display(x)

    // Output: 
    // http://www.mrap.org/upload/XYZ234/status/UP-AB2345
}

func ExampleIsAlive() {
    e := createEndpoint("https://www.catfur.org")
    x := e.IsAlive()
    display(x)

    // Output: https://www.catfur.org/isAlive
}

func ExampleSave() {
    e := createEndpoint("https://dcc.icgc.org:8080")
    studyId := "XYZ234"
    uploadId := "UP-AB2345"

    x := e.Save(studyId, uploadId, true)
    display(x)

    y := e.Save(studyId, uploadId, false)
    display(y)

    // Output: 
    // https://dcc.icgc.org:8080/upload/XYZ234/save/UP-AB2345?ignoreAnalysisIdCollisions=true
    // https://dcc.icgc.org:8080/upload/XYZ234/save/UP-AB2345?ignoreAnalysisIdCollisions=false
}

func ExamplePublish() {
    e := createEndpoint("http://example.org:12345")
    studyId, analysisId := "XQA-𝜆123","A2345-999-7012"
    x := e.Publish(studyId, analysisId)
    display(x)

   // Output: http://example.org:12345/studies/XQA-%F0%9D%9C%86123/analysis/publish/A2345-999-7012
}

func ExampleSuppress() {
    e := createEndpoint("http://www.testing.com")
    studyId, analysisId := "ABC123", "AN-123579"
    x := e.Suppress(studyId, analysisId)
    display(x)

    // Output: http://www.testing.com/studies/ABC123/analysis/suppress/AN-123579

}

func ExampleGetAnalysis() {
   e := createEndpoint("http://abc.de")
   studyId, analysisId := "ABC123", "AN-123579"
   x := e.GetAnalysis(studyId, analysisId)
   display(x)
  // Output: http://abc.de/studies/ABC123/analysis/AN-123579

}

func ExampleGetAnalysisFiles() {
   e := createEndpoint("https://localhost:8080")
   studyId, analysisId := "XYZ2345","13"
   x := e.GetAnalysisFiles(studyId, analysisId)
   display(x)
  // Output: https://localhost:8080/studies/XYZ2345/analysis/13/files
}

func ExampleIdSearch() {
   e := createEndpoint("http://abc.de:123")
   studyId := "ABC123"
   x := e.IdSearch(studyId)
   display(x)
   // Output: http://abc.de:123/studies/ABC123/analysis/search/id
}

func ExampleInfoSearch() {

   e := createEndpoint("http://xyz.ai:23")
   studyId := "XYZ2345"
   x := e.InfoSearch(studyId)
   display(x)
   // Output: http://xyz.ai:23/studies/XYZ2345/analysis/search/info
}
