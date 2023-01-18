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

function sectionReachTopScreen(dom_element) {
  if (dom_element === undefined) {
      return false;
  }

  const element_rect = dom_element.getBoundingClientRect();

  const element_top = element_rect.top;

  const viewport_height = Math.max(document.documentElement.clientHeight, window.innerHeight || 0);

  return element_top <= viewport_height * .14;
}

export function checkLastComponentCollition(sections_color_schemas, check_every=50) {
    this.skiped_sections = this.skiped_sections !== undefined ? this.skiped_sections : check_every;
    if (this.skiped_sections < check_every) {
        this.skiped_sections++;
        return;
    }

    this.skiped_sections = 0;

    const sections = Object.values(sections_color_schemas);

    if (sections.length === 0) {
        return;
    }

    let last_collided_section = sections[0];
    
    for (let section of sections) {
        if (sectionReachTopScreen(section.ref)) {
            if (section.ord > last_collided_section.ord) {
                last_collided_section = section;
            }
        }
    }

    last_collided_section.color_schema.define();
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

