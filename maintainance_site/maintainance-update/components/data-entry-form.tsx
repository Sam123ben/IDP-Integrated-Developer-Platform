// components/data-entry-form.tsx
'use client';

import dynamic from 'next/dynamic';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { API_BASE_URL } from '@/config';
import { toast } from 'sonner';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Form, FormField, FormItem, FormLabel, FormControl } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { baseSchema } from '@/app/types/form-schemas';

const SystemComponentsForm = dynamic(() => import('./forms/system-components-form').then(mod => mod.SystemComponentsForm), { ssr: false });
const MaintenanceWindowsForm = dynamic(() => import('./forms/maintenance-windows-form').then(mod => mod.MaintenanceWindowsForm), { ssr: false });
const MaintenanceUpdatesForm = dynamic(() => import('./forms/maintenance-updates-form').then(mod => mod.MaintenanceUpdatesForm), { ssr: false });

export function DataEntryForm() {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formData, setFormData] = useState<any>(null); // Stores specific form data

  const form = useForm({
    resolver: zodResolver(baseSchema),
    defaultValues: {
      email: '',
      updateType: 'system_components', // Default to system components
    },
  });

  const updateType = form.watch('updateType');

  const onSubmit = async (data: any) => {
    setIsSubmitting(true);
    try {
      let endpoint = '';
      let payload = {};

      switch (data.updateType) {
        case 'maintenance_windows':
          payload = { ...formData }; // Includes maintenance window form data
          endpoint = '/maintenance-windows';
          break;
        case 'maintenance_updates':
          payload = {
            message: formData.message,
            maintenance_window_id: formData.maintenance_window_id
          };
          endpoint = '/maintenance-updates';
          break;
        default:
          throw new Error('Invalid update type');
      }

      const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorMessage = await response.text();
        throw new Error(errorMessage || 'Failed to submit data');
      }

      toast.success('Successfully submitted!', {
        description: 'Your changes have been saved to the database.',
      });

      // Reset the form
      form.reset({
        ...form.getValues(),
        email: data.email,
      });
      setFormData(null); // Clear specific form data
    } catch (error: any) {
      toast.error(error.message || 'Failed to submit form');
    } finally {
      setIsSubmitting(false);
    }
  };

  const renderFormByType = () => {
    switch (updateType) {
      case 'system_components':
        return <SystemComponentsForm />;
      case 'maintenance_windows':
        return <MaintenanceWindowsForm onFormDataChange={setFormData} />;
      case 'maintenance_updates':
        return <MaintenanceUpdatesForm onFormDataChange={setFormData} />;
      default:
        return null;
    }
  };

  return (
    <div className="container mx-auto p-6">
      <Card className="p-8 shadow-2xl">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            {/* Email Field */}
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder="your.email@company.com" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />

            {/* Update Type Selector */}
            <FormField
              control={form.control}
              name="updateType"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Update Type</FormLabel>
                  <Select onValueChange={field.onChange} value={field.value}>
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select update type" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="system_components">System Components</SelectItem>
                      <SelectItem value="maintenance_windows">Maintenance Windows</SelectItem>
                      <SelectItem value="maintenance_updates">Maintenance Updates</SelectItem>
                    </SelectContent>
                  </Select>
                </FormItem>
              )}
            />

            {/* Form for Specific Update Type */}
            <div>{renderFormByType()}</div>

            {/* Conditional Submit Button */}
            {updateType !== 'system_components' && (
              <div className="flex justify-center">
                <Button type="submit" disabled={isSubmitting}>
                  {isSubmitting ? 'Submitting...' : 'Submit'}
                </Button>
              </div>
            )}
          </form>
        </Form>
      </Card>
    </div>
  );
}
