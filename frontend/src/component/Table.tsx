import React, { useContext } from 'react';
import { DataContext } from '../DataContext';

const Table: React.FC = () => {
  const { data } = useContext(DataContext);

  console.log(data.processedWrongs)

  return (
    <div>
      <h1>What you have to do</h1>
      <div className="flex flex-col">
        <div className="-m-1.5 overflow-x-auto">
          <div className="p-1.5 min-w-full inline-block align-middle">
            <div className="border border-gray-400 overflow-hidden rounded">
              <table className="min-w-full divide-y divide-gray-400">
                <thead>
                  <tr>
                    <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Name</th>
                    <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Simirarity</th>
                    <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Simirarity RATE</th>
                    <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Link</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-500">
                  {data && data.recommends && data.processedWrongs && data.similarities && data.recommends.map((recommend, index) => {
                    const taskUrl = `https://atcoder.jp/contests/${recommend.split(' ')[0].toLowerCase()}/tasks/${recommend.replace(/-/g, '_').replace(/\s+/g, '').toLowerCase()}`;
                    return (
                      <tr key={index}>
                        <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-800">{recommend}</td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{data && data.processedWrongs && data.processedWrongs[index]}</td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{data && data.similarities && data.similarities[index]}</td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-sm font-medium">
                          <p>
                            <a className="text-blue-600 hover:underline hover:decoration-blue-600" href={taskUrl} target="_blank" rel="noopener noreferrer">
                              Learn
                            </a>
                          </p>
                        </td>
                      </tr>
                    );
                  })}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Table;
