import * as Yup from "yup";

export function configurationValues() {
  return {
    email: "",
    code: "",
  };
}

export function validationSchema() {
  return Yup.object().shape({
    email: Yup.string().email(true).required(true),
    code: Yup.string().required(true),
  });
}
