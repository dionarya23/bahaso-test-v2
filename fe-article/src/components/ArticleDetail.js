import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

const ArticleDetails = () => {
  const { id } = useParams();
  const [article, setArticle] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchArticle = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article?id=${id}`);
        if (!response.ok) {
          throw new Error('Could not fetch the article');
        }
        const data = await response.json();
        setArticle(data.data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchArticle();
  }, [id]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!article) return <div>No article found.</div>;

  const createMarkup = html => ({ __html: html });

  const formatDate = dateString => new Date(dateString).toLocaleDateString("en-US", {
    weekday: 'long', year: 'numeric', month: 'long', day: 'numeric'
  });

  return (
    <div className="container mt-3">
      <h1>{article.title}</h1>
      {article.image_url && <img src={article.image_url} alt={article.title} className="img-fluid" />}
      <p><small>Published on: {formatDate(article.created_at)}</small></p>
      <div dangerouslySetInnerHTML={createMarkup(article.content)}></div>
    </div>
  );
};

export default ArticleDetails;
