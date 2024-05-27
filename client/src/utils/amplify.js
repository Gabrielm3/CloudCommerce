import { Amplify } from "aws-amplify"

export function configureAmplify() {
    Amplify.configure({
        Auth: {
          Cognito: {
            userPoolClientId: 'abcdefghij1234567890',
            userPoolId: 'us-east-1_abcd1234',
            loginWith: { // Optional
              oauth: {
                domain: 'abcdefghij1234567890-29051e27.auth.us-east-1.amazoncognito.com',
                scopes: ['openid','email','phone','profile','aws.cognito.signin.user.admin'],
                redirectSignIn: ['http://localhost:3000/','https://example.com/'],
                redirectSignOut: ['http://localhost:3000/','https://example.com/'],
                responseType: 'code',
              },
              username: 'true',
              email: 'false', // Optional
              phone: 'false', // Optional
            }
          }
        }
      });
}
