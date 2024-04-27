package com.spring3react.crudpractice.passByValue;

public class Employee{

    private String name;
    private double salary;

    Employee(String n){
        this.name = n;
    }

    public String getName(){
        return this.name;
    }

    public double getSalary(){
        return this.salary;
    }

    public void setSalary(double s){
        this.salary = s;
    }

    public String toString(){
        return "Employee: name: " + this.getName() + ", salary: "+ this.getSalary();
    }

}