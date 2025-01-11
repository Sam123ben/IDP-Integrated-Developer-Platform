// src/components/ContactInfo.tsx
import React from 'react';
import { Mail, Twitter } from 'lucide-react';

export function ContactInfo() {
  return (
    <div className="space-y-6">
      <h2 className="text-xl font-semibold text-gray-900">Need to get in touch?</h2>
      <div className="grid md:grid-cols-2 gap-4">
        <a
          href="mailto:support@example.com"
          className="flex items-center space-x-3 p-4 rounded-lg border border-gray-200 hover:border-indigo-300 hover:bg-indigo-50 transition-colors duration-200"
        >
          <Mail className="w-5 h-5 text-indigo-600" />
          <span className="text-gray-700">support@example.com</span>
        </a>
      </div>
    </div>
  );
}