import React, { createContext, useState, ReactNode } from 'react';
import internal from 'stream';

// データの型定義
interface DataType {
  submissions?: number[];
  performances?: number[];
  rates?: number[];
  contestNames?: string[];
  recommends?: string[];
  similarities?: number[];
  processedWrongs?: string[];
  UserIdentifier?: number;
  // 他のデータフィールドもここに追加
}

// デフォルト値の型注釈を追加
const defaultContextValue = {
  data: {} as DataType, // 空のオブジェクトで初期化
  setData: (data: DataType) => {}
};

// デフォルト値をcreateContextに渡す
export const DataContext = createContext(defaultContextValue);

// Provider コンポーネントの定義
export const DataProvider = ({ children }: { children: ReactNode }) => {
  const [data, setData] = useState<DataType>({}); // DataType を使用

  return (
      <DataContext.Provider value={{ data, setData }}>
          {children}
      </DataContext.Provider>
  );
};
