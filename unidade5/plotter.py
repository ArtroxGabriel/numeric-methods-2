import subprocess
import matplotlib.pyplot as plt


def run_go_and_plot():
    """
    Executes a Go program, captures its output, and plots the data.
    """
    try:
        result = subprocess.run(
            ["go", "run", "main.go"], capture_output=True, text=True, check=True
        )

        # The standard output from the Go program is stored in result.stdout.
        go_output = result.stdout

        # Process the output: split into lines and convert each line to a float.
        # We use strip() to remove leading/trailing whitespace and filter out empty lines.
        lines = go_output.strip().split("\n")
        data_points = [float(line) for line in lines if line]

        # Plot the data using matplotlib.
        plt.figure(figsize=(12, 7))
        plt.plot(data_points, marker="o", linestyle="-", color="b")

        # Add titles and labels for clarity.
        plt.title("Data Plot from Go Program Output", fontsize=16)
        plt.xlabel("Data Point Index", fontsize=12)
        plt.ylabel("Value", fontsize=12)
        plt.grid(True, which="both", linestyle="--", linewidth=0.5)
        plt.xticks(
            range(len(data_points))
        )  # Ensure all points have a tick on the x-axis

        # Save the plot to a file.
        plot_filename = "go_output_plot.png"
        plt.savefig(plot_filename)

        print(f"Successfully executed Go program and saved plot to '{plot_filename}'")
        print("\nCaptured Data:")
        for i, val in enumerate(data_points):
            print(f"  Point {i}: {val}")

    except FileNotFoundError:
        print("Error: The 'go' command was not found.")
        print(
            "Please ensure that Go is installed and that its binary is in your system's PATH."
        )
    except subprocess.CalledProcessError as e:
        print(f"Error executing Go program: {e}")
        print(f"Stderr: {e.stderr}")
    except ValueError as e:
        print(f"Error parsing the output from the Go program: {e}")
        print("Please ensure the Go program outputs only numbers, one per line.")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")


if __name__ == "__main__":
    run_go_and_plot()
