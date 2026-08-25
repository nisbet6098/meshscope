import { useEffect, useState } from 'react';

function App() {
  const [services, setServices] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch('/api/v1/services')
      .then((res) => {
        if (!res.ok) throw new Error('Failed to reach MeshScope Core API');
        return res.json();
      })
      .then((data) => {
        setServices(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  return (
    <div style={{ fontFamily: 'system-ui, sans-serif', padding: '2rem', maxWidth: '800px', margin: '0 auto' }}>
      <h1>🔍 MeshScope Control Panel</h1>
      <p style={{ color: '#666' }}>Zero-Touch Local Service Discovery</p>
      
      <hr style={{ margin: '1.5rem 0', borderColor: '#eee' }} />

      <h2>Active Services</h2>

      {loading && <p>Scanning local nodes...</p>}
      {error && <p style={{ color: 'red' }}>Error connecting to backend: {error}</p>}

      {!loading && !error && (
        <div style={{ display: 'grid', gap: '1rem', marginTop: '1rem' }}>
          {services.map((srv) => (
            <div key={srv.id} style={{ border: '1px solid #ddd', borderRadius: '8px', padding: '1rem', background: '#fafafa' }}>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <h3 style={{ margin: 0 }}>{srv.name}</h3>
                <span style={{ background: '#e6fffa', color: '#047857', padding: '0.25rem 0.5rem', borderRadius: '4px', fontSize: '0.875rem' }}>
                  {srv.status}
                </span>
              </div>
              <p style={{ margin: '0.5rem 0 0', color: '#555', fontFamily: 'monospace' }}>
                {srv.ip}:{srv.port}
              </p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default App;