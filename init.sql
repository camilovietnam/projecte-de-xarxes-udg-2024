-- init.sql

-- Create database
CREATE DATABASE "projecte-de-xarxes";

-- Connect to the newly created database
\c projecte-de-xarxes;

-- Create table
CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);
