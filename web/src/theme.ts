import {createMuiTheme, responsiveFontSizes} from "@material-ui/core/styles"

export const theme = responsiveFontSizes(
  createMuiTheme({
    overrides: {
      MuiCssBaseline: {
        "@global": {
          body: {
            backgroundColor: "#f8f7f3",
          },
        },
      },
      MuiSelect: {
        select: {
          backgroundColor: "#ffffff",
          fontSize: "14px",
          padding: "8px",
          "&:focus": {
            borderRadius: 30,
            backgroundColor: "#ffffff",
          },
        },
      },
      MuiOutlinedInput: {
        root: {
          "& $notchedOutline": {
            borderStyle: "none",
            boxShadow: "0px 2px 10px 1px #100d0d0d",
          },
          "&$focused $notchedOutline": {
            borderStyle: "none",
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
