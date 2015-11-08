var util = require('util');
var _ = require('lodash');
var exec = require('child_process').exec;
var child;
require('shelljs/global');

module.exports = function (grunt) {
  var cwd = __dirname;
  var watchdirs = [
    cwd + '/../**/*.go',
    cwd + '/../**/*.sh'
  ];
  grunt.initConfig({
    watch   : {
      scripts: {
        files  : watchdirs,
        tasks  : ['vect_build'],
        options: {
          debounceDelay: 600
        }
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.registerTask('vect_build', function () {

    var vect = {dir: cwd + '/../', name: 'vect'};

    var dirs = [vect];

    for (var i = 0; i < dirs.length; ++i) {
      var directory = dirs[i].dir;
      cd(directory); //cd to dir
      var cmd = (directory + 'build.sh');
      var name = dirs[i].name;
      child = exec(cmd, exec_callback.bind(null, name));
      cd(__dirname);//cd out to curdir
    }

    function exec_callback(name, error, stdout, stderr) {
      var line = '..........................' + name + '............................';
      console.log(line);
      if (error !== null) {
        console.log('exec error: ' + error);
        console.log('stderr: ' + stderr);
      }
      else {
        console.log('stdout: ' + stdout);

      }
    }
  });

  grunt.registerTask('default', ['watch']);
  grunt.task.run(['vect_build', 'watch']);

};