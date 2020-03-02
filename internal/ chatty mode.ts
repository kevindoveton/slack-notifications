// chatty mode
// ----------------
// show by default
// blacklist
// whitelist

// standard mode
// ----------------
// inherit all rules from chatty
// blacklist
// whitelist

// focus mode
// ----------------
// ignore by default
// whitelist

interface Modes {
  [mode: string]: {
    [id: string]: {
      rules: {
        after?: string // '14:55'
        before?: string // '12:32'
        day?: string // mon | tues | wed | thurs | fri | sat | sun | mon-fri etc
      }[]
    }
  },
}

interface DeclaredModes {
  modes: {
    name: string
    inherits?: string
    order: ('show'|'hide'|'blacklist'|'whitelist')[]
  }[]
}

type Rules = DeclaredModes & Modes;

const rules: Rules = {
  modes: [{name: 'chatty', order: ['show', 'blacklist', 'whitelist']}],
  // chatty: {
  //   '@chilli': {
  //     rules: [{}]
  //   }
  // }
};