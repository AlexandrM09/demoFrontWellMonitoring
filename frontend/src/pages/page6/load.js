
import Papa from 'papaparse'
export function getParselet () {
  // Papa.DefaultDelimiter = ','
  return Papa.parse('static/2A022.csv')
}
