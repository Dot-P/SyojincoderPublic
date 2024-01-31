import React, { useState, useEffect } from 'react';

const AlgoTable: React.FC = () => {

  interface Problem {
    problem_id: string;
    difficulties: string;
    classifier: string;
    url: string;
  }

  const [problems, setProblems] = useState<Problem[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/algo_search");
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setProblems(data);
      } catch (error) {
        console.error('Fetch error:', error);
      }
    };

    fetchData();
  }, []);

return (
  <div className="flex flex-col">
    <div className="-m-1.5 overflow-x-auto">
      <div className="p-1.5 min-w-full inline-block align-middle">
        <div className="border rounded-lg divide-y divide-gray-200">
          <div className="overflow-hidden">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Name</th>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Difficulty</th>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Category</th>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Link</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {problems.map((problem, index) => (
                  <tr key={index}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{problem.problem_id}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{problem.difficulties}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{problem.classifier}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-start text-sm font-medium">
                      <a href={problem.url} className="text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none">Learn</a>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
  );
};

export default AlgoTable;