import React from 'react';
import { Link } from 'react-router-dom';

const ArticleList = ({ articles, editable = false, onDelete }) => {

  const renderContent = (content) => {
    return { __html: content };
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  };

  return (
    <div className="container mt-4">
      <h1 className="mb-4">Articles</h1>
      {articles.length > 0 ? (
        <div className="row">
          {articles.map(article => (
            <div key={article.id} className="col-md-4 mb-3">
              <div className="card h-100">
                {article.image_url && (
                  <img
                    src={article.image_url}
                    alt={article.title}
                    className="card-img-top"
                    style={{ height: '200px', objectFit: 'cover' }}
                  />
                )}
                <div className="card-body">
                  <h5 className="card-title">{article.title}</h5>
                  
                  <div className="card-text" dangerouslySetInnerHTML={renderContent(article.content)}></div>
                  
                  <p className="card-text">
                    <small className="text-muted">Published on {formatDate(article.created_at)}</small>
                  </p>
                </div>
                <div className="card-footer d-flex justify-content-between">
                  <Link to={`/articles/${article.id}`} className="btn btn-primary">Read More</Link>
                  {(editable && article.is_owned) && (
                    <>
                      <Link to={`/edit-article/${article.id}`} className="btn btn-secondary">Edit</Link>
                      <button className="btn btn-danger" onClick={() => onDelete(article.id)}>Delete</button>
                    </>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <p>No articles found.</p>
      )}
    </div>
  );
};

export default ArticleList;
