
import SwiftUI

struct MeView: View {
    @State var name: String
    @State var id: String
    @State var stOrTe: String
    @State var login = true
    
    var body: some View {
        VStack{
            if self.login == false {
                ContentView()
            }else{
                VStack {
                    VStack(alignment: .leading) {
                        Form {
                            Text("　　名前 ： \(self.name)")
                            Text("　　身元 ： \(self.stOrTe)")
                            Text("　　ID　 ： \(self.id)")
                            Text("")
                        }
                        .offset(x: -7, y: -40) //underlineを左右に揃える
                        .frame(width: 220, height: 169) //中に移す
                        .clipped() //背景を切り取る
                    }
                }
            }
        }
    }
}

struct MeView_Previews: PreviewProvider {
    static var previews: some View {
        MeView(name: "name??", id: "id??", stOrTe: "stOrTe??")
    }
}
