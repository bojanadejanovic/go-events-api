import { writable } from 'svelte/store';
import { browser } from '$app/environment';

interface User {
    id: number;
    email: string;
}

export const user = writable<User | null>(null);
export const token = writable<string | null>(null);

export function setAuth(userData: User, authToken: string) {
    user.set(userData);
    token.set(authToken);
    if (browser) {
        localStorage.setItem('token', authToken);
    }
}

export function clearAuth() {
    user.set(null);
    token.set(null);
    if (browser) {
        localStorage.removeItem('token');
    }
}

// Initialize from localStorage only in browser
if (browser) {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
        token.set(storedToken);
    }
} 