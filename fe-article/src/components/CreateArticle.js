import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';

function CreateArticle() {
  const [article, setArticle] = useState({
    title: '',
    content: '',
    image_url: '' 
  });
  const [image, setImage] = useState(null);
  const [uploading, setUploading] = useState(false);
  const navigate = useNavigate();

  const handleTitleChange = (e) => {
    const { name, value } = e.target;
    setArticle(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleContentChange = (value) => {
    setArticle(prev => ({
      ...prev,
      content: value
    }));
  };

  const handleImageChange = (e) => {
    setImage(e.target.files[0]);
  };

  const handleImageUpload = async () => {
    if (!image) return;
  
    setUploading(true);
  
    const formData = new FormData();
    formData.append('file', image);
    formData.append('api_key', process.env.REACT_APP_CLOUDINARY_API_KEY);
    formData.append('timestamp', Math.round((new Date()).getTime() / 1000));
    formData.append('upload_preset', process.env.REACT_APP_CLOUDINARY_UPLOAD_PRESET);
  
    try {
      const response = await fetch(process.env.REACT_APP_CLOUDINARY_API_URL, {
        method: 'POST',
        body: formData,
      });
  
      const data = await response.json();
  
      if (data.secure_url) {
        setArticle(prev => ({ ...prev, image_url: data.secure_url }));
      } else {
        throw new Error('Image upload failed');
      }
  
      setUploading(false);
    } catch (error) {
      console.error('Error uploading image:', error);
      setUploading(false);
    }
  };
  

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (!article.image_url) {
        alert('Please upload an image before submitting.');
        return;
      }

      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `${JSON.parse(localStorage.getItem("user")).accessToken}`,
        },
        body: JSON.stringify(article),
      });

      if (!response.ok) {
        throw new Error('Failed to create the article');
      }

      navigate('/admin');
    } catch (error) {
      console.error('Error creating article:', error);
    }
  };

  return (
    <div className="container mt-5">
      <h2>Create New Article</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="title" className="form-label">Title</label>
          <input
            type="text"
            className="form-control"
            id="title"
            name="title"
            value={article.title}
            onChange={handleTitleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="content" className="form-label">Content</label>
          <ReactQuill 
            theme="snow"
            value={article.content}
            onChange={handleContentChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="image" className="form-label">Upload Image</label>
          <input
            type="file"
            className="form-control"
            id="image"
            onChange={handleImageChange}
            accept="image/*"
          />
          <button type="button" className="btn btn-info mt-2" onClick={handleImageUpload} disabled={uploading}>
            {uploading ? 'Uploading...' : 'Upload Image'}
          </button>
        </div>
        {article.image_url && <img src={article.image_url} alt="Uploaded" style={{ width: '100%', maxHeight: '200px', objectFit: 'cover' }} />}
        <button type="submit" className="btn btn-primary mt-3">Submit</button>
      </form>
    </div>
  );
}

export default CreateArticle;