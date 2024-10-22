components {
  id: "animalController"
  component: "/main/scripts/Controller1.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"cat\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/test.atlas\"\n"
  "}\n"
  ""
  position {
    y: 54.0
    z: 0.5
  }
  scale {
    x: 0.404244
    y: 0.506549
  }
}
