package data

// This willl be responsible for taking loaded data  from the text files (in the form of a map of strings), and turning
// those into entities, as required. For example, we may have loaded several potion definitions into memory from the
// definition files, and we now want to use those in the games. In order to do that, we would find the potion we want
// to load, and then take that definition and turn it into an entity in the ECS. We can do this as many times as we
// need per potion definition. In this way, we have an easy way of loading data file information into the ECS.
