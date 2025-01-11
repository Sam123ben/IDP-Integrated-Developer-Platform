// components/forms/maintenance-updates-form.tsx
'use client';

import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';

type MaintenanceUpdateInputs = {
  message: string;
  maintenance_window_id: number;
};

export function MaintenanceUpdatesForm({
  onFormDataChange,
}: {
  onFormDataChange: (data: MaintenanceUpdateInputs) => void;
}) {
  const { control, watch, getValues } = useForm<MaintenanceUpdateInputs>({
    defaultValues: {
      message: '',
      maintenance_window_id: 1, // Default value, adjust as needed
    },
  });

  // Watch for changes and update parent
  useEffect(() => {
    const subscription = watch((value, { name, type }) => {
      if (type === 'change') {
        // Get current form values
        const currentValues = getValues();
        // Pass to parent component
        onFormDataChange({
          message: currentValues.message,
          maintenance_window_id: currentValues.maintenance_window_id
        });
      }
    });

    return () => subscription.unsubscribe();
  }, [watch, onFormDataChange, getValues]);

  return (
    <div className="space-y-6">
      <FormField
        control={control}
        name="message"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Message</FormLabel>
            <FormControl>
              <Input 
                placeholder="Enter update message..." 
                {...field}
                onChange={(e) => {
                  field.onChange(e);
                  // Ensure the value is immediately available
                  onFormDataChange({
                    ...getValues(),
                    message: e.target.value
                  });
                }}
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />

      <FormField
        control={control}
        name="maintenance_window_id"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Maintenance Window ID</FormLabel>
            <FormControl>
              <Input 
                type="number" 
                {...field}
                onChange={(e) => {
                  const value = parseInt(e.target.value);
                  field.onChange(value);
                  onFormDataChange({
                    ...getValues(),
                    maintenance_window_id: value
                  });
                }}
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
    </div>
  );
}