package weapon


type Weapon struct {
  Damage  float32
}


func NewWeapon() *Weapon {
  return &Weapon{
    Damage: 5,
  }
}


