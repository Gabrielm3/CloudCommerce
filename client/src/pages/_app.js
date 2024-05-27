import 'semantic-ui-css/semantic.min.css'
import { configureAmplify } from "@/utils/amplify"
import { AuthContext } from "@/contexts"

configureAmplify();

export default function App(props) {
  const { Component, pageProps } = props;

  return (
    <AuthProvider>
      <Component {...pageProps} />
    </AuthProvider>
  )
}
