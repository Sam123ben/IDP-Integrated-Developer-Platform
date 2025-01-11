// components/forms/maintenance-windows-form.tsx
'use client';

import { useEffect } from 'react';
import { useForm, Controller } from 'react-hook-form';
import ReactDatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';

type MaintenanceWindowFormInputs = {
  start_time: Date | null;
  estimated_duration: number;
  description: string;
  issue_fixed: boolean;
};

export function MaintenanceWindowsForm({
  onFormDataChange,
}: {
  onFormDataChange: (data: MaintenanceWindowFormInputs) => void;
}) {
  const { control, watch, getValues } = useForm<MaintenanceWindowFormInputs>({
    defaultValues: {
      start_time: null,
      estimated_duration: 0,
      description: '',
      issue_fixed: false,
    },
  });

  useEffect(() => {
    const subscription = watch((value, { name, type }) => {
      if (type === 'change') {
        const currentValues = getValues();
        onFormDataChange(currentValues);
      }
    });

    return () => subscription.unsubscribe();
  }, [watch, onFormDataChange, getValues]);

  return (
    <div className="space-y-6">
      <FormField
        control={control}
        name="start_time"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Start Time</FormLabel>
            <div className="flex gap-2 items-center">
              <FormControl className="flex-1">
                <Controller
                  name="start_time"
                  control={control}
                  render={({ field }) => (
                    <ReactDatePicker
                      selected={field.value}
                      onChange={(date) => field.onChange(date)}
                      showTimeSelect
                      timeFormat="HH:mm aa"
                      timeIntervals={15}
                      dateFormat="MMMM d, yyyy h:mm aa"
                      className="w-full border rounded-md px-2 py-2"
                      placeholderText="Select a date and time"
                    />
                  )}
                />
              </FormControl>
              {field.value && (
                <button
                  type="button"
                  onClick={() => field.onChange(null)}
                  className="text-sm text-blue-600 hover:text-blue-800 hover:underline"
                >
                  Clear
                </button>
              )}
            </div>
            <FormMessage />
          </FormItem>
        )}
      />

      <FormField
        control={control}
        name="estimated_duration"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Estimated Duration (minutes)</FormLabel>
            <FormControl>
              <Input 
                type="number" 
                placeholder="e.g., 240" 
                {...field} 
                onChange={(e) => field.onChange(Number(e.target.value))}
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />

      <FormField
        control={control}
        name="description"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Description</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Describe the maintenance" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />

      <FormField
        control={control}
        name="issue_fixed"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Issue Fixed</FormLabel>
            <Select
              onValueChange={(value) => field.onChange(value === 'true')}
              defaultValue={field.value ? 'true' : 'false'}
            >
              <FormControl>
                <SelectTrigger>
                  <SelectValue placeholder="Select status" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectItem value="false">No</SelectItem>
                <SelectItem value="true">Yes</SelectItem>
              </SelectContent>
            </Select>
            <FormMessage />
          </FormItem>
        )}
      />
    </div>
  );
}