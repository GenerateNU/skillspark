import axios, {
  AxiosRequestConfig,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from 'axios';

interface CustomAxiosRequestConfig extends InternalAxiosRequestConfig {
  _retry?: boolean;
}

// Platform-agnostic storage helper
const getStorageItem = (key: string): string | null => {
  if (typeof window !== 'undefined' && window.localStorage) {
    return localStorage.getItem(key);
  }
  return null;
};

const removeStorageItem = (key: string): void => {
  if (typeof window !== 'undefined' && window.localStorage) {
    localStorage.removeItem(key);
  }
};

const getBaseURL = () => {
  if (typeof process !== 'undefined' && process.env) {
    if (process.env.EXPO_PUBLIC_API_BASE_URL) {
      return process.env.EXPO_PUBLIC_API_BASE_URL;
    }
    if (process.env.VITE_API_BASE_URL) {
      return process.env.VITE_API_BASE_URL;
    }
    if (process.env.NEXT_PUBLIC_API_BASE_URL) {
      return process.env.NEXT_PUBLIC_API_BASE_URL;
    }
  }
  
  return 'http://localhost:8080';
};

const apiClient = axios.create({
  baseURL: getBaseURL(),
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

apiClient.interceptors.request.use(
  (config) => {
    const token = getStorageItem('temp_jwt') || getStorageItem('jwt');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

let isRetrying = false;

const handleLogout = () => {
  removeStorageItem('jwt');
  removeStorageItem('userId');
  removeStorageItem('recentlyViewedStudents');
  
  if (typeof document !== 'undefined') {
    document.cookie.split(';').forEach((cookie) => {
      const eqPos = cookie.indexOf('=');
      const name = eqPos > -1 ? cookie.substring(0, eqPos).trim() : cookie.trim();
      document.cookie = `${name}=;expires=Thu, 01 Jan 1970 00:00:00 GMT;path=/`;
    });
    
    if (typeof window !== 'undefined') {
      window.location.href = '/login';
    }
  }
};

apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const status = error.response?.status;
    const config = error.config as CustomAxiosRequestConfig;

    if (status === 401) {
      if (!isRetrying && !config._retry) {
        isRetrying = true;
        config._retry = true;

        await new Promise((resolve) => setTimeout(resolve, 100));

        try {
          const result = await apiClient.request(config);
          isRetrying = false;
          return result;
        } catch (retryError) {
          isRetrying = false;
          console.error('Unauthorized access');
          handleLogout();
          return Promise.reject(retryError);
        }
      } else {
        console.error('Unauthorized access');
        handleLogout();
      }
    } else if (status === 403) {
      console.error('Forbidden access');
    } else if (status === 404) {
      console.error('Resource not found');
    } else if (status >= 500) {
      console.error('Server error occurred');
    } else {
      console.error('An error occurred:', error.message);
    }

    return Promise.reject(error);
  }
);

export function customInstance<T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig
): Promise<T> {
  return apiClient({
    ...config,
    ...options,
    withCredentials: true,
  }).then((response: AxiosResponse<T>) => response.data);
}

export default apiClient;