import {createMuiTheme, responsiveFontSizes} from "@material-ui/core/styles"

export const theme = responsiveFontSizes(
  createMuiTheme({
    overrides: {
      MuiSelect: {
        select: {
          fontSize: "14px",
          padding: "8px",
          "&:focus": {
            borderRadius: 30,
          },
        },
      },
    },
    palette: {
      primary: {
        main: "#000000",
      },
      secondary: {
        main: "#ffffff",
      },
    },
    shape: {
      borderRadius: 30,
    },
    typography: {
      fontFamily: [
        "-apple-system",
        "BlinkMacSystemFont",
        '"Segoe UI"',
        "Roboto",
        '"Helvetica Neue"',
        "Arial",
        "sans-serif",
        '"Apple Color Emoji"',
        '"Segoe UI Emoji"',
        '"Segoe UI Symbol"',
      ].join(","),
      fontSize: 16,
      h1: {
        fontWeight: 900,
      },
      h2: {
        fontWeight: 700,
      },
      h3: {
        fontWeight: 700,
      },
      h4: {
        fontWeight: 700,
      },
      h5: {
        fontWeight: 700,
      },
      h6: {
        fontWeight: 700,
      },
    },
  }),
  {factor: 4},
)
