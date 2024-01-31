import pandas as pd
import time
import os
from openai import OpenAI

# ファイルの読み込み
df = pd.read_csv('problem.csv')

# APIキーを設定
api_key = 'YOUR_OPENAI_API_KEY'

client = OpenAI(api_key=api_key)

# Embeddingを取得する関数を定義
def get_embedding(text, model="text-embedding-ada-002"):
    response = client.embeddings.create(model=model, input=text)
    embedding_data = response.data[0].embedding
    return embedding_data

# 開始位置を設定
start = 1100  # ここを変更して任意の開始位置を指定

# 処理対象の行数を設定
num_rows_to_process = 100

# Embeddingデータを格納するリストを作成
embeddings = []

count = 0

# dfの特定の範囲に対してget_embeddingを呼び出す
for _, row in df.iloc[start:start + num_rows_to_process].iterrows():
    embedding = get_embedding(row['Content'], model='text-embedding-ada-002')
    # 'Category'と'Name'の値を辞書に追加
    embeddings.append({
        'Category': row['Category'], 
        'Name': row['Name'], 
        'Content': row['Content'], 
        'embedding': embedding
    })
    
    count += 1
    print("process...",count,"%")

    time.sleep(25)  # 25秒間の遅延

# 新しいデータフレームを作成
new_df = pd.DataFrame(embeddings)

# CSVファイルが存在しない場合は作成し、存在する場合は追加
csv_file = 'embedding_data.csv'
if not os.path.exists(csv_file):
    new_df.to_csv(csv_file, index=False)
else:
    with open(csv_file, 'a') as f:
        new_df.to_csv(f, header=False, index=False)

