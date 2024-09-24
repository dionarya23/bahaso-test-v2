import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import ReactQuill from 'react-quill'; 
import 'react-quill/dist/quill.snow.css';

function EditArticle() {
  const { id } = useParams(); 
  const navigate = useNavigate();
  const [article, setArticle] = useState({
    title: '',
    content: '',
    image_url: '',
  });
  const [image, setImage] = useState(null); 
  const [uploading, setUploading] = useState(false); 
  const [error, setError] = useState(null); 

  
  useEffect(() => {
    const fetchArticle = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article/admin?id=${id}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${JSON.parse(localStorage.getItem("user")).accessToken}`,
          }
        });

        if (!response.ok) {
          throw new Error('Failed to fetch the article');
        }

        const data = await response.json();
        setArticle(data.data);
      } catch (error) {
        console.error('Error fetching article:', error);
        setError('Failed to load article');
      }
    };

    fetchArticle();
  }, [id]);

  
  const handleTitleChange = (e) => {
    const { name, value } = e.target;
    setArticle(prev => ({ ...prev, [name]: value }));
  };

  
  const handleContentChange = (value) => {
    setArticle(prev => ({ ...prev, content: value }));
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
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${JSON.parse(localStorage.getItem("user")).accessToken}`,
        },
        body: JSON.stringify(article),
      });

      if (!response.ok) {
        throw new Error('Failed to update the article');
      }

      navigate('/admin');
    } catch (error) {
      console.error('Error updating article:', error);
      setError('Failed to update article');
    }
  };

  return (
    <div className="container mt-5">
      <h2>Edit Article</h2>

      {error && <div className="alert alert-danger">{error}</div>}

      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="title" className="form-label">Title</label>
          <input
            type="text"
            className="form-control"
            id="title"
            name="title"
            value={article.title || ''}
            onChange={handleTitleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label htmlFor="content" className="form-label">Content</label>
          <ReactQuill 
            theme="snow"
            value={article.content || ''}
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
          <button
            type="button"
            className="btn btn-info mt-2"
            onClick={handleImageUpload}
            disabled={uploading}
          >
            {uploading ? 'Uploading...' : 'Upload Image'}
          </button>
        </div>

        {article.image_url && (
          <img
            src={article.image_url}
            alt="Uploaded"
            style={{ width: '100%', maxHeight: '200px', objectFit: 'cover' }}
          />
        )}

        <button type="submit" className="btn btn-primary mt-3">Update Article</button>
      </form>
    </div>
  );
}

export default EditArticle;
