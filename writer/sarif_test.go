package writer

import "os"

func ExampleSarif() {
	w, _ := NewSarif(os.Stdout, SarifOption{ToolName: "super-linter"})
	for _, e := range testErrs {
		w.Write(e)
	}
	w.Flush()
	// Output:
	// {
	//   "$schema": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
	//   "runs": [
	//     {
	//       "results": [
	//         {
	//           "level": "warning",
	//           "locations": [
	//             {
	//               "physicalLocation": {
	//                 "artifactLocation": {
	//                   "uri": "path/to/file1",
	//                   "uriBaseId": "%SRCROOT%"
	//                 },
	//                 "region": {
	//                   "endColumn": 15,
	//                   "endLine": 2,
	//                   "startColumn": 14,
	//                   "startLine": 1
	//                 }
	//               }
	//             }
	//           ],
	//           "message": {
	//             "text": "hello"
	//           }
	//         },
	//         {
	//           "level": "note",
	//           "locations": [
	//             {
	//               "physicalLocation": {
	//                 "artifactLocation": {
	//                   "uri": "path/to/file1",
	//                   "uriBaseId": "%SRCROOT%"
	//                 },
	//                 "region": {
	//                   "startColumn": 14,
	//                   "startLine": 2
	//                 }
	//               }
	//             }
	//           ],
	//           "message": {
	//             "text": "vim"
	//           }
	//         },
	//         {
	//           "level": "error",
	//           "locations": [
	//             {
	//               "physicalLocation": {
	//                 "artifactLocation": {
	//                   "uri": "file2",
	//                   "uriBaseId": "%SRCROOT%"
	//                 },
	//                 "region": {
	//                   "startColumn": 14,
	//                   "startLine": 2
	//                 }
	//               }
	//             }
	//           ],
	//           "message": {
	//             "text": "emacs"
	//           }
	//         },
	//         {
	//           "level": "error",
	//           "locations": [
	//             {
	//               "physicalLocation": {
	//                 "artifactLocation": {
	//                   "uri": "file2",
	//                   "uriBaseId": "%SRCROOT%"
	//                 },
	//                 "region": {
	//                   "startColumn": 1,
	//                   "startLine": 14
	//                 }
	//               }
	//             }
	//           ],
	//           "message": {
	//             "text": "neovim"
	//           }
	//         }
	//       ],
	//       "tool": {
	//         "driver": {
	//           "name": "super-linter"
	//         }
	//       }
	//     }
	//   ],
	//   "version": "2.1.0"
	// }
}
