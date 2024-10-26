import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

# Load the CSV file
csv_file = 'results.csv'  # Change this to your output file name if needed
data = pd.read_csv(csv_file)

# Check if your CSV file has the correct structure
# If it contains only the average reward, we need a different approach
# Here we assume each agent's results are stored by episode

# Prepare a DataFrame suitable for line plotting
# Assuming data has 'Algorithm', 'Episodes', 'Reward'
# If your data structure is different, adjust accordingly
data['Episodes'] = data['Episodes'].astype(int)  # Ensure 'Episodes' is treated as an integer

# Set the aesthetics for the plots
sns.set(style="whitegrid")

# Create a line plot
plt.figure(figsize=(10, 6))
line_plot = sns.lineplot(data=data, x='Episodes', y='Reward', hue='Algorithm', marker='o')

# Set titles and labels
plt.title('Agent Reward Over Episodes', fontsize=16)
plt.xlabel('Number of Episodes', fontsize=14)
plt.ylabel('Average Reward', fontsize=14)
plt.legend(title='Algorithm')

# Show the plot
plt.tight_layout()
plt.savefig('agent_reward_over_episodes.png')  # Save the plot as a PNG file
plt.show()
