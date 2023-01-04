const NAVBAR_CSS_ATTRIBUTES = {
    'color': 'navbar-schema-color',
    'background': 'navbar-schema-background',
    'border': 'navbar-schema-border',
    'contrast': 'navbar-schema-contrast'
}

function defineNavbarColorSchema() {
    document.documentElement.style.setProperty(`--${NAVBAR_CSS_ATTRIBUTES.color}`, this.color);
    document.documentElement.style.setProperty(`--${NAVBAR_CSS_ATTRIBUTES.background}`, this.background);
    document.documentElement.style.setProperty(`--${NAVBAR_CSS_ATTRIBUTES.border}`, this.border);
    document.documentElement.style.setProperty(`--${NAVBAR_CSS_ATTRIBUTES.contrast}`, this.contrast);
}

class ColorSchema {
  constructor(color, contrast) {
    this.color = color;
    this.contrast = contrast;
    this.background = contrast;
    this.border = color;
  }

  define = () => {
    throw new Error('ColorSchema.define() is not implemented');
  }
}

export const supported_components = {
    NAVBAR: 'navbar'
}

export default (schema, component) => {
    const color_schema = new ColorSchema(schema.color, schema.contrast);

    switch (component) {
        case 'navbar':
            color_schema.define = defineNavbarColorSchema.bind(color_schema);
            break;
        default:
            color_schema = null;
            break;
    }

    return color_schema;
} 

