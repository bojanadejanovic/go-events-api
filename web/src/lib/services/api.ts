import { token } from '$lib/stores/auth';
import { get } from 'svelte/store';
import { browser } from '$app/environment';

const API_URL = 'http://localhost:8080';

interface Event {
    id: number;
    name: string;
    description: string;
    date_time: string;
    location: string;
}

interface LoginResponse {
    token: string;
    user: {
        id: number;
        email: string;
    };
}

export async function login(email: string, password: string): Promise<LoginResponse> {
    const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    });

    if (!response.ok) {
        throw new Error('Login failed');
    }

    return response.json();
}

export async function signup(email: string, password: string): Promise<LoginResponse> {
    const response = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    });

    if (!response.ok) {
        throw new Error('Signup failed');
    }

    return response.json();
}

function getAuthHeaders(): Record<string, string> {
    const authToken = get(token);
    if (!authToken) return {};
    return {
        'Authorization': `Bearer ${authToken}`
    };
}

export async function getEvents(): Promise<Event[]> {
    const response = await fetch(`${API_URL}/events`, {
        headers: {
            'Content-Type': 'application/json',
            ...getAuthHeaders()
        }
    });

    if (!response.ok) {
        throw new Error('Failed to fetch events');
    }

    return response.json();
}

export async function registerForEvent(eventId: number): Promise<void> {
    const response = await fetch(`${API_URL}/events/${eventId}/register`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            ...getAuthHeaders()
        }
    });

    if (!response.ok) {
        throw new Error('Failed to register for event');
    }
}

export async function cancelRegistration(eventId: number): Promise<void> {
    const response = await fetch(`${API_URL}/events/${eventId}/register`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            ...getAuthHeaders()
        }
    });

    if (!response.ok) {
        throw new Error('Failed to cancel registration');
    }
}

export async function getRegisteredEvents(): Promise<Event[]> {
    const response = await fetch(`${API_URL}/events/registered`, {
        headers: {
            'Content-Type': 'application/json',
            ...getAuthHeaders()
        }
    });

    if (!response.ok) {
        throw new Error('Failed to fetch registered events');
    }

    return response.json();
} 