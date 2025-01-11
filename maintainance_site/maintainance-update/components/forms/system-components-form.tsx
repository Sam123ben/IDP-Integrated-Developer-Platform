// components/forms/system-components-form.tsx
'use client';

import { useState, useEffect } from 'react';
import { useFormContext } from 'react-hook-form';
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { API_BASE_URL } from '@/config';
import { toast } from 'sonner';
import { CheckCircleIcon } from 'lucide-react';

export function SystemComponentsForm() {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { control, getValues, setValue, reset } = useFormContext();

  useEffect(() => {
    // Set the default value for componentStatus if not explicitly set
    const currentStatus = getValues('componentStatus');
    if (!currentStatus) {
      setValue('componentStatus', 'operational');
    }
  }, [getValues, setValue]);

  const handleAction = async (method: 'POST' | 'PUT') => {
    setIsSubmitting(true);
    const values = getValues();

    // Transform the data to match the API expectations
    const payload = {
      name: values.componentName,
      type: values.componentType,
      status: values.componentStatus || 'operational', // Default to 'operational' if undefined
    };

    try {
      const endpoint = method === 'POST' ? '/system-components' : '/system-components/update';

      const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        method,
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorData = await response.text();
        throw new Error(errorData || 'Failed to submit data');
      }

      const data = await response.json();

      toast.success(`${method === 'POST' ? 'Created' : 'Updated'} successfully`, {
        description: data.message || 'Your changes have been saved to the database.',
        icon: <CheckCircleIcon className="h-4 w-4 text-green-500" />,
      });

      if (method === 'POST') {
        reset({
          ...values,
          componentName: '',
          componentType: '',
          componentStatus: 'operational', // Reset to 'operational'
        });
      }
    } catch (error: any) {
      console.error('Error:', error);
      toast.error(error.message || 'Failed to submit data');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="space-y-8">
      <FormField
        control={control}
        name="componentName"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Component Name</FormLabel>
            <FormControl>
              <Input placeholder="e.g., Database Server" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <FormField
        control={control}
        name="componentType"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Component Type</FormLabel>
            <FormControl>
              <Input placeholder="e.g., Database" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <FormField
        control={control}
        name="componentStatus"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Component Status</FormLabel>
            <Select
              onValueChange={field.onChange}
              value={field.value || 'operational'} // Default to 'operational'
            >
              <FormControl>
                <SelectTrigger>
                  <SelectValue placeholder="Select status" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectItem value="operational">Operational</SelectItem>
                <SelectItem value="degraded">Degraded</SelectItem>
                <SelectItem value="maintenance">Maintenance</SelectItem>
              </SelectContent>
            </Select>
            <FormMessage />
          </FormItem>
        )}
      />

      <div className="flex gap-4 justify-center mt-6">
        <Button
          type="button"
          className="bg-blue-500 hover:bg-blue-600 text-white"
          onClick={() => handleAction('POST')}
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Creating...' : 'Create'}
        </Button>
        <Button
          type="button"
          className="bg-green-500 hover:bg-green-600 text-white"
          onClick={() => handleAction('PUT')}
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Updating...' : 'Update'}
        </Button>
      </div>
    </div>
  );
}