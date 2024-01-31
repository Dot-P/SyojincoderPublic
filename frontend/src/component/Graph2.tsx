import React, { useContext, useEffect, useState } from 'react';
import Chart from 'react-apexcharts';
import { DataContext } from '../DataContext';

// シリーズデータの型を定義
interface SeriesData {
  name: string;
  data: number[];
}

const Graph2: React.FC = () => {
  const { data } = useContext(DataContext); // DataContext からデータを取得

  const [options, setOptions] = useState({
    chart: {
      id: 'apexchart-example'
    },
    xaxis: {
      categories: ["8週間前", "7週間前", "6週間前", "5週間前", "4週間前", "3週間前", "2週間前", "1週間前", "0週間前"]
    }
  });

  const [series, setSeries] = useState<SeriesData[]>([{
    name: 'Submission Num',
    data: []
  }]);

  useEffect(() => {
    if (data && data.submissions) {
      setSeries([{
        name: 'Submission Num',
        data: data.submissions
      }]);
    }
  }, [data]); // data の変更を監視

  return (
    <div>
      <h1>Your Submissions</h1>
      <Chart options={options} series={series} type="bar" width={500} height={350} />
    </div>
  );
};

export default Graph2;
