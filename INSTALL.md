# Recipe Stats Calculator CLI

## Installation

1. Pull the code.
2. Run `docker compose up`.
3. Go into the golang container by running:
   ```
   docker exec -it container_id /bin/bash
   ```

4. Run the following command:
   ```
   go run calculator.go -file=hf_test_calculation_fixtures.json -recipes="Apple,Cake,Pasta" -postcode="10224" -from=10AM -to=7PM
   ```
   (The provided json file is too large, please download manually and add to root path)
## Arguments

- `file`: Path to custom fixtures file (path to file).
- `recipes`: Custom recipe names to search by (string separated with commas).
- `postcode`: Custom postcode.
- `from`: Start time for the delivery search window(10AM).
- `to`: End time for the delivery search window(3PM).

## Test

To run the tests, navigate to the `./stats` directory and execute the following command:
```
go test
```

## Clarifications

- **Parsing JSON**: The time complexity of parsing JSON data is typically O(n), where n is the size of the input data.

- **Counting unique recipe names**: The script iterates over all recipes once to count the unique recipe names. The time complexity is O(n), where n is the number of recipes.

- **Sorting recipe names**: The `sort.Slice` function internally uses the quicksort algorithm, which has an average time complexity of O(n log n).

- **Finding the busiest postcode**: The script iterates over all recipes once to count the occurrences of each postcode. The time complexity is O(n), where n is the number of recipes.

- **Counting deliveries within a time range**: The script iterates over all recipes once to count the deliveries within a time range for a specific postcode. The time complexity is O(n), where n is the number of recipes.

- **Listing recipe names containing specific words**: The script iterates over all recipes once to check if each recipe contains the specified words. The time complexity is O(n), where n is the number of recipes.

Overall, the time complexity of the script is O(n log n).


## Improvement

Since the max number of distinct recipe names is known(lower than 2k),  replacing `sort.Slice` with counting sort can reduce Overall time complexity to  O(n)
