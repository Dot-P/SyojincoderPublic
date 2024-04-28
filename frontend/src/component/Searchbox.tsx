import React, { useState, useEffect, useContext } from 'react';
import { DataContext } from '../DataContext';

const SearchBox: React.FC<{ setUserName: (userName: string) => void }> = ({ setUserName }) => {
    const { setData } = useContext(DataContext); // DataContext から setData を取得
    const [localUserName, setLocalUserName] = useState('');
    const [isValid, setValid] = useState(0);
    const [timer, setTimer] = useState<number | null>(null);
    const { data } = useContext(DataContext);

    useEffect(() => {
      if (data.UserIdentifier === 1) {
          setValid(1);
      } else if (data.UserIdentifier === 2) {
          setValid(2);
      }
  }, [data.UserIdentifier]);

    const handleSubmit = async () => {
        try {
            const response = await fetch("https://syojincoder-3d33c5ad4332.herokuapp.com/dashboard", {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ userName: localUserName })
            });
            const data = await response.json();
            setData(data); // 取得したデータを DataContext 経由で共有
        } catch (error) {
            console.error('Error:', error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      setLocalUserName(e.target.value); // ローカルの userName を更新
    };
  
    useEffect(() => {
      if (timer) {
        clearTimeout(timer);
      }
      if (localUserName !== '') {
        const newTimer = window.setTimeout(() => {
          setUserName(localUserName); // 2秒後に親コンポーネントの userName を更新
          handleSubmit();
        }, 2000);
        setTimer(newTimer as unknown as number);
      }
  
      return () => {
        if (timer) {
          clearTimeout(timer);
        }
      };
    }, [localUserName]);

    return (
      <div>
        <input
          type="text"
          className="py-3 px-4 block w-1/4 border-gray-500 rounded-lg text-sm focus:border-blue-400 focus:ring-blue-400 disabled:opacity-50 disabled:pointer-events-none"
          placeholder="User Name"
          value={localUserName}
          onChange={handleInputChange}
        />
        {
          isValid === 0 ? null : 
          isValid === 1 ? <p className="text-sm text-red-600 mt-2" id="hs-validation-name-error-helper">Please enter a valid User Name.</p> :
          <p className="text-sm text-teal-600 mt-2" id="hs-validation-name-success-helper">Accept!</p>
        }
        <br></br>
      </div>
    );
};

export default SearchBox;

