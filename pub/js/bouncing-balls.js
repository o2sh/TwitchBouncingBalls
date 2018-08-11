(function () {
    var width = window.innerWidth - 100;
    var height = window.innerHeight - 100;
    var scene = new THREE.Scene();
    camera = new THREE.OrthographicCamera(-width / 2, width / 2, height / 2, -height / 2, 1, 1000);
    camera.position.set(0, 0, 30)

    var renderer = new THREE.WebGLRenderer();
    renderer.setClearColor(0x52f97c);
    renderer.setSize(window.innerWidth, window.innerHeight);
    document.getElementById('render-area').appendChild(renderer.domElement);

    window.addEventListener('resize', function () {
        var width = window.innerWidth;
        var height = window.innerHeight;
        renderer.setSize(width, height);
        camera.aspect = width / height;
        camera.updateProjectionMatrix();
    })

    var spheres = []
    var sphereRadius = 30
    var total = `{{.}}`
    console.log(total)
    for (var i = 0; i < total; i++) {
        var geometry = new THREE.SphereGeometry(sphereRadius, 10, 10);
        var material = new THREE.MeshBasicMaterial({
            color: 0x00a8a5,
            wireframe: false
        });

        spheres[i] = new sphere(new THREE.Mesh(geometry, material));

        scene.add(spheres[i].mesh);
    }


    function sphere(mesh) {
        var edge = [width / 2, height / 2]
        this.mesh = mesh;
        this.direction = [
            Math.round(Math.random()) == 1 ? 1 : -1,
            Math.round(Math.random()) == 1 ? 1 : -1,
        ];

        this.speed = Math.random() * 5 + 5; // Pixels per millisecond

        this.updatePosition = function () {
            this.mesh.position.x += this.direction[0] * this.speed;
            this.mesh.position.y += this.direction[1] * this.speed;
        }

        this.updateDirection = function () {
            if (Math.abs(this.mesh.position.x) + sphereRadius >= edge[0])
                this.direction[0] = -this.direction[0];
            if (Math.abs(this.mesh.position.y) + sphereRadius >= edge[1])
                this.direction[1] = -this.direction[1];
        }

    }

    function update() {
        for (var i = 0; i < spheres.length; i++) {
            spheres[i].updatePosition();
            spheres[i].updateDirection();
        }
    }

    function render() {
        renderer.render(scene, camera);
    }

    function gameLoop() {
        window.animationId = requestAnimationFrame(gameLoop);
        render();
        update();
    }
    window.onload = gameLoop();

})();