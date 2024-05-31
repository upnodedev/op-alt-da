package common

// TxDataVersion1 is the version number for batcher transactions containing
// plasma commitments. It should not collide with DerivationVersion which is still
// used downstream when parsing the frames.
const TxDataVersion1 = 1
