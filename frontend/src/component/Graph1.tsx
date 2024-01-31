import React, { useContext, useEffect, useState } from 'react';
import Chart from 'react-apexcharts';
import { ApexOptions } from 'apexcharts';
import { DataContext } from '../DataContext';

interface Graph1Props {
  userName?: string;
}

interface SeriesData {
  name: string;
  data: number[];
}

const Graph1: React.FC<Graph1Props> = ({ userName }) => {
  const { data } = useContext(DataContext); // DataContext からデータを取得
  
  const [options, setOptions] = useState<ApexOptions>({
    chart: {
      height: 350,
      type: 'area'
    },
    dataLabels: {
      enabled: false
    },
    stroke: {
      curve: 'smooth'
    },
    xaxis: {
      type: 'category',
      categories: [] // カテゴリを空に初期化
    },
    tooltip: {
      x: {
        format: 'dd/MM/yy HH:mm'
      },
    }
  });

  const [series, setSeries] = useState<SeriesData[]>([
    { name: 'Your Performance', data: [] },
    { name: 'Your Rate', data: [] }
  ]);  

  useEffect(() => {
    // ユーザ名とデータの存在を確認
    if (data && data.performances && data.rates) {
      // オプションとシリーズデータを更新
      setOptions(prevOptions => ({
        ...prevOptions,
        xaxis: {
          ...prevOptions.xaxis,
          categories: data.contestNames
        }
      }));

      setSeries([
        { name: 'Your Performance', data: data.performances },
        { name: 'Your Rate', data: data.rates }
      ]);
    }
  }, [userName, data]); // userName と data の変更を監視

  return (
    <div>
      <h1>Your Performance and Rates</h1>
      <Chart options={options} series={series} type="area" width="180%" height={350} />
    </div>
  );
};

export default Graph1;
