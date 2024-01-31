import pandas as pd

# 既存のCSVファイルを読み込む
embedding_df = pd.read_csv('embedding_data.csv')

# 元のCSVファイルを読み込む
original_df = pd.read_csv('problem.csv')

# 最初の100行のみを取得
original_df = original_df.iloc[:100]

# 両方のDataFrameを結合（マージ）
# 'Content'カラムをキーとして使用
merged_df = pd.merge(original_df[['Category', 'Name', 'Content']], embedding_df, on='Content')

# 新しいCSVファイルとして保存
merged_df.to_csv('embedding_data.csv', index=False)
