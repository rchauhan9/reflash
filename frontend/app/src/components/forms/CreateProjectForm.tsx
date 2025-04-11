import React, { useState, DragEvent, ChangeEvent } from 'react';
import { useFormik, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { CardContent } from '@/components/ui/card';
import EmojiPicker, { EmojiClickData } from 'emoji-picker-react';
import { useCreateProject } from '@/hooks/api/use-study';

interface FormValues {
  name: string;
  icon: string;
  description: string;
  files: File[];
}

const validationSchema = Yup.object({
  name: Yup.string().required('Name is required'),
  icon: Yup.string().required('Icon is required'),
  description: Yup.string().required('Description is required'),
  files: Yup.array().min(0, ''),
});

const CreateProjectForm: React.FC = () => {
  const [showEmojiPicker, setShowEmojiPicker] = useState<boolean>(false);
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const today = new Date();
  const currentMonthYear = today.toLocaleString('default', { month: 'long', year: 'numeric' });

  const createProject = useCreateProject()

  const formik = useFormik<FormValues>({
    initialValues: {
      name: '',
      icon: '',
      description: '',
      files: [],
    },
    validationSchema,
    onSubmit: (values: FormValues, actions: FormikHelpers<FormValues>) => {
      console.log('Form values:', values);
      const project = { name: values.name, icon: values.icon, description: values.description }
      console.log("Project:", project);
      createProject.mutate(project)
    },
  });

  const handleEmojiClick = (emojiObject: EmojiClickData): void => {
    formik.setFieldValue('icon', emojiObject.emoji);
    setShowEmojiPicker(false);
  };

  const MAX_FILE_SIZE_MB = 10;

  const handleDrop = (event: DragEvent<HTMLDivElement>): void => {
    event.preventDefault();
    const filesArray = Array.from(event.dataTransfer.files).filter(
      (file) => file.size <= MAX_FILE_SIZE_MB * 1024 * 1024,
    );
    const newFiles = [...uploadedFiles, ...filesArray];
    setUploadedFiles(newFiles);
    formik.setFieldValue('files', newFiles);
  };

  const handleRemoveFile = (index: number): void => {
    const newFiles = uploadedFiles.filter((_, i) => i !== index);
    setUploadedFiles(newFiles);
    formik.setFieldValue('files', newFiles);
  };

  return (
    <CardContent>
      <form onSubmit={formik.handleSubmit} className='space-y-4'>
        {/* Name and Icon Row */}
        <div className='flex items-end space-x-4'>
          <div className='flex-1'>
            <Label htmlFor='name'>Name</Label>
            <Input
              id='name'
              name='name'
              placeholder='e.g. Advanced Biology'
              onChange={formik.handleChange}
              value={formik.values.name}
            />
            {formik.touched.name && formik.errors.name && (
              <p className='text-red-500 text-sm'>{formik.errors.name}</p>
            )}
          </div>
          <div className='mb-1'>
            <Label className='sr-only' htmlFor='icon'>
              Icon
            </Label>
            <Button
              id='icon'
              type='button'
              variant='outline'
              onClick={() => setShowEmojiPicker(!showEmojiPicker)}
            >
              {formik.values.icon ? formik.values.icon : 'Add icon'}
            </Button>
            {showEmojiPicker && (
              <div className='mt-2 z-10 absolute'>
                <EmojiPicker onEmojiClick={handleEmojiClick} />
              </div>
            )}
          </div>
        </div>

        {/* Description */}
        <div>
          <Label htmlFor='description'>Description</Label>
          <Input
            id='description'
            name='description'
            placeholder={`e.g. Study plan for ${currentMonthYear}`}
            onChange={formik.handleChange}
            value={formik.values.description}
          />
          {formik.touched.description && formik.errors.description && (
            <p className='text-red-500 text-sm'>{formik.errors.description}</p>
          )}
        </div>

        {/* File Upload with Drag and Drop */}
        <div>
          <Label htmlFor='files'>Upload Files</Label>
          <div
            onDrop={handleDrop}
            onDragOver={(e) => e.preventDefault()}
            className='border border-dashed border-gray-400 p-4 rounded-md text-center cursor-pointer hover:bg-gray-50'
          >
            Drag and drop files here or click to upload
            <Input
              id='files'
              name='files'
              type='file'
              multiple
              className='hidden'
              onChange={(event: ChangeEvent<HTMLInputElement>) => {
                const filesArray = Array.from(event.currentTarget.files || []).filter(
                  (file) => file.size <= MAX_FILE_SIZE_MB * 1024 * 1024,
                );
                const newFiles = [...uploadedFiles, ...filesArray];
                setUploadedFiles(newFiles);
                formik.setFieldValue('files', newFiles);
              }}
            />
          </div>
          {formik.touched.files && formik.errors.files && (
            <p className='text-red-500 text-sm mt-1'>{formik.errors.files}</p>
          )}
          <ul className='mt-2 space-y-1'>
            {uploadedFiles.map((file, index) => (
              <li
                key={index}
                className='flex justify-between items-center text-sm bg-gray-100 rounded px-3 py-1'
              >
                <span>{file.name}</span>
                <button
                  type='button'
                  onClick={() => handleRemoveFile(index)}
                  className='text-red-500 hover:text-red-700 text-xs'
                >
                  ✕
                </button>
              </li>
            ))}
          </ul>
        </div>

        <div className='flex justify-end'>
          <Button type='submit' disabled={createProject.isLoading}>{createProject.isLoading ? 'Submitting...' : 'Submit'}</Button>
        </div>
      </form>
    </CardContent>
  );
};

export default CreateProjectForm;
