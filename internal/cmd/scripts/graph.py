### This script processes heart rate data from a CSV file, calculates daily and hourly averages,
### and generates plots for visualization. It also exports the plots as HTML files.

### Vibe Coded with Gemini 2.5 Flash
### Tyler Christensen


import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import plotly.express as px
import plotly.io as pio
import os
import argparse


parser = argparse.ArgumentParser(
    description="Process heart rate data from a CSV file and generate plots."
)
parser.add_argument(
    "--csv", type=str, default="./export/records.csv", help="Path to the CSV file."
)
parser.add_argument(
    "--export",
    type=str,
    default="./graphs",
    help="Path to export the plots as HTML files.",
)
args = parser.parse_args()
csv_file_path = args.csv
export_path = args.export

if export_path and not os.path.exists(export_path):
    export_path = os.path.abspath(export_path)
    print(f"Creating export directory at: {export_path}")
    os.makedirs(export_path)

try:
    df = pd.read_csv(
        csv_file_path,
        parse_dates=["DateTime"],
        date_format="%m/%d/%y %H:%M:%S",
    )
except Exception as e:
    print(f"Error loading CSV: {e}")
    print("Trying without date_format and inferring...")
    df = pd.read_csv(csv_file_path, parse_dates=["DateTime"])

print(f"Successfully loaded {len(df)} rows.")
print(df.head())
print(df.info())

df["BPM"] = pd.to_numeric(df["BPM"], errors="coerce")
df.dropna(subset=["BPM"], inplace=True)

df["date_only"] = df["DateTime"].dt.date
daily_avg_bpm = df.groupby("date_only")["BPM"].mean().reset_index()
daily_avg_bpm["date_only"] = pd.to_datetime(daily_avg_bpm["date_only"])

print("\nDaily Average BPM:")
print(daily_avg_bpm.head())

df["hour_of_day"] = df["DateTime"].dt.floor("h")
hourly_avg_bpm = df.groupby("hour_of_day")["BPM"].mean().reset_index()
print("\nHourly Average BPM (first 5):")
print(hourly_avg_bpm.head())


plt.figure(figsize=(15, 7))
sns.lineplot(data=daily_avg_bpm, x="date_only", y="BPM")
plt.title("Daily Average Heart Rate Over Time")
plt.xlabel("Date")
plt.ylabel("Average BPM")
plt.grid(True)
plt.tight_layout()
plt.savefig(os.path.join(export_path, "daily_avg_bpm.png"))


last_day_in_data = df["DateTime"].dt.date.max()
one_day_df = df[df["DateTime"].dt.date == last_day_in_data].copy()

if not one_day_df.empty:
    plt.figure(figsize=(15, 7))
    sns.lineplot(data=one_day_df, x="DateTime", y="BPM")
    plt.title(f"Heart Rate for {last_day_in_data} (Raw Data)")
    plt.xlabel("Date and Time")
    plt.ylabel("BPM")
    plt.grid(True)
    plt.tight_layout()
    plt.savefig(os.path.join(export_path, f"raw_data_{last_day_in_data}.png"))
else:
    print(f"No data found for {last_day_in_data} to plot raw data.")

fig_daily = px.line(
    daily_avg_bpm,
    x="date_only",
    y="BPM",
    title="Interactive Daily Average Heart Rate Over Time",
)

pio.write_html(
    fig_daily,
    os.path.join(export_path, "daily_avg_bpm.html"),
    auto_open=True,
)

fig_hourly = px.line(
    hourly_avg_bpm,
    x="hour_of_day",
    y="BPM",
    title="Interactive Hourly Average Heart Rate Over Time",
)

pio.write_html(
    fig_hourly,
    os.path.join(export_path, "hourly_avg_bpm.html"),
    auto_open=True,
)
