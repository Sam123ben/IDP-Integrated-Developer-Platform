// app/components/cors-test.tsx
'use client';

import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';

const CorsTest = () => {
  const [response, setResponse] = useState('');
  const [error, setError] = useState('');

  const makeRequest = async (method: 'PUT' | 'POST') => {
    setResponse('');
    setError('');
    
    const payload = {
      name: "Database Server",
      type: "Database",
      status: "operational"
    };

    try {
      console.log(`Making ${method} request to /api/system-components`);
      
      const response = await fetch('http://localhost:8080/api/system-components', {
        method,
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
        body: JSON.stringify(payload)
      });

      const data = await response.text();
      console.log(`${method} Response:`, data);
      setResponse(`${method} Response: ${data}`);
      
    } catch (err) {
      console.error(`${method} Error:`, err);
      setError(`${method} Error: ${err.message}`);
    }
  };

  return (
    <Card className="p-6 space-y-4">
      <div className="space-x-4">
        <Button 
          onClick={() => makeRequest('PUT')} 
          className="bg-blue-500 hover:bg-blue-600 text-white"
        >
          Test PUT Request
        </Button>
        <Button 
          onClick={() => makeRequest('POST')} 
          className="bg-green-500 hover:bg-green-600 text-white"
        >
          Test POST Request
        </Button>
      </div>
      
      {response && (
        <div className="p-4 bg-green-100 text-green-800 rounded">
          {response}
        </div>
      )}
      
      {error && (
        <div className="p-4 bg-red-100 text-red-800 rounded">
          {error}
        </div>
      )}
    </Card>
  );
};

export default CorsTest;