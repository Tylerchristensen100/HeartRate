# HeartRate: Visualize Your Fitbit Heart Rate Data

Visualize and analyze your **heart rate data** exported from Fitbit via [Google Takeout](https://takeout.google.com/). This tool processes your raw data and generates insightful graphs.

---

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

* **Python 3**: This tool relies on Python for data processing and graph generation. You can download it from [python.org](https://www.python.org/downloads/).

### Installation

No complex installation is required. Simply download the `heartRate` executable and the associated Python script.

---

## 🏃‍♀️ Usage
To run the application, navigate to its directory in your terminal and execute:

`./heartRate [options]`
### Options
Here are the available command-line options:

- import string: Required. This specifies the path to the directory containing your raw heart rate recordings from Google Takeout.

      Default: ./data

      Example: -import "/path/to/your/fitbit_takeout/Heart Rate"

- export string: This determines the directory where processed data files (like CSVs) will be saved.

      Default: ./export

- export-name string: Use this to set the base name for the exported processed files (e.g., if you set it to "my_records", you'll get my_records.csv).

      Default: records

- graph-export string: This is the path to the directory where the generated graphs (e.g., PNG images) will be saved by the Python script.

      Default: ./graphs

#### Example
Let's say your Fitbit heart rate data is located in ~/Downloads/Takeout/Fitbit/Heart Rate. You want to export processed data to a processed_data folder and graphs to a my_heart_graphs folder, with exported files named my_heart_records. You'd run:


      `$ ./heartRate -import "~/Downloads/Takeout/Fitbit/Heart Rate" 
           -export "./processed_data"`