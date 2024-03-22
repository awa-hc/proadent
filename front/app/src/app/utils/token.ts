function getTokenPayload(token: string): any {
  try {
    const payload = token.split('.')[1]; // Obtén el payload del token
    const decodedPayload = atob(payload.replace(/-/g, '+').replace(/_/g, '/')); // Decodifica de Base64 URL
    return JSON.parse(decodedPayload); // Parsea el JSON
  } catch (error) {
    return null;
  }
}

export default function GetUserIdFromToken(token: string): number | null {
  const payload = getTokenPayload(token);
  if (payload) {
    const userId =
      payload[
        'http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier'
      ];
    const userIdint = parseInt(userId, 10);
    return userIdint;
  }
  return null;
}

export function GetRoleFromToken(token: string): string | null {
  const payload = getTokenPayload(token);
  if (payload) {
    const role =
      payload['http://schemas.microsoft.com/ws/2008/06/identity/claims/role'];
    return role;
  }
  return null;
}
//   }
