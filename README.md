##Grupo32-Laboratorio-1
- Juan Francisco Baquedano  202073611-1
- Vicente Figueroa          202073627-8
- Vicente Ruiz              202004642-5

Le asignamos los siguientes servidores a las máquinas virtuales:
- Dist 125: central y asia
- Dist 126: europa
- Dist 127: américa
- Dist 128: oceania y rabbit

Como ejecutar el programa:
- Primero clonar el github del Grupo32
- Hacer cd Grupo32-Laboratorio-1
- Para ejecutar los archivos hacer primero el make docker-rabbit,
después make docker-[region] (en [region] poner la region correspondiente),
y por último make docker-central. Por lo cual se necesitan 6 consolas abiertas.

Posdata:
- Tuvimos problemas con la ejecución del programa en las máquinas virtuales
ya que nos aparecia un problema de conexión rechazada al conectarnos al rabbit,
por lo cual algunas veces nos servia el programa y otras veces no, el problema
se refería a que algunos puertos estaban cerrados por ende tuvimos que abrir 
los puertos manualmente.
